package network.platon.utils;

import com.example.contract.Solc;

import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.util.concurrent.Semaphore;
import java.util.concurrent.TimeUnit;

/**
 * @title CompileUtil
 * @description 编译工具类
 * @author qcxiao
 * @updateTime 2019/12/27 14:39
 */
public class CompileUtil {
    private final Semaphore permit = new Semaphore(100, true);

    public void evmCompile(String file, String buildPath) throws Exception {
        try {
            permit.acquire();
            Solc.compile(file, buildPath);
        } catch (Exception e) {
            e.printStackTrace();
            throw new Exception(e);
        } finally {
            permit.release();
        }
    }

    public void wasmCompile(String file, String buildPath) throws Exception {
        try {
            permit.acquire();
            // /usr/local/bin/platon-cpp
            String[] args = new String[]{"/bin/bash", "/usr/local/bin/platon-cpp", "-o", file, buildPath, "-abigen"};
            Process ps = Runtime.getRuntime().exec(args);
            ps.waitFor(2, TimeUnit.SECONDS);
            BufferedReader br = new BufferedReader(new InputStreamReader(ps.getInputStream()));
            StringBuffer sb = new StringBuffer();

            String line;
            while((line = br.readLine()) != null) {
                sb.append(line).append("\n");
            }

            String result = sb.toString();
        } catch (Exception e) {
            e.printStackTrace();
            throw new Exception(e);
        } finally {
            permit.release();
        }
    }
}
