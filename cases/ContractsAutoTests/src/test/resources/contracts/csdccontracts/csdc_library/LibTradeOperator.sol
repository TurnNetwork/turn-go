pragma solidity ^0.4.12;
/**
*@file      LibTradeOperator.sol
*@author    yiyating
*@time      2017-07-03
*@desc      the defination of LibTradeOperator
*/

import "../utillib/LibInt.sol";
import "../utillib/LibString.sol";
import "../utillib/LibJson.sol";
import "../utillib/LibStack.sol";

library LibTradeOperator {
    using LibInt for *;
    using LibString for *;
    using LibJson for *;
    using LibTradeOperator for *;

    struct TradeOperator {
        address id;
        uint brokerId;          //券商id
        string name;            //姓名
        string department;      //部门名称
        string phone;           //电话
        string fax;             //传真
        string mobile;          //手机
        string email;           //邮箱
    }

    /**
    *@desc fromJson for TradeOperator
    *      Generated by juzhen SolidityStructTool automatically.
    *      Not to edit this code manually.
    */
    function fromJson(TradeOperator storage _self, string _json) internal returns(bool succ) {
        _self.reset();
        if(LibJson.push(_json) == 0) {
            return false;
        }

        if (!_json.isJson()) {
            LibJson.pop();
            return false;
        }

        _self.id = _json.jsonRead("id").toAddress();
        _self.brokerId = _json.jsonRead("brokerId").toUint();
        _self.name = _json.jsonRead("name");
        _self.department = _json.jsonRead("department");
        _self.phone = _json.jsonRead("phone");
        _self.fax = _json.jsonRead("fax");
        _self.mobile = _json.jsonRead("mobile");
        _self.email = _json.jsonRead("email");
        
        LibJson.pop();
        return true;
    }

    /**
    *@desc toJson for TradeOperator
    *      Generated by juzhen SolidityStructTool automatically.
    *      Not to edit this code manually.
    */
    function toJson(TradeOperator storage _self) internal constant returns (string _json) {
        uint len = 0;
        len = LibStack.push("{");
        len = LibStack.appendKeyValue("id", _self.id);
        len = LibStack.appendKeyValue("brokerId", _self.brokerId);
        len = LibStack.appendKeyValue("name", _self.name);
        len = LibStack.appendKeyValue("department", _self.department);
        len = LibStack.appendKeyValue("phone", _self.phone);
        len = LibStack.appendKeyValue("fax", _self.fax);
        len = LibStack.appendKeyValue("mobile", _self.mobile);
        len = LibStack.appendKeyValue("email", _self.email);
        len = LibStack.append("}");
        _json = LibStack.popex(len);
    }

    /**
    *@desc update for TradeOperator
    *      Generated by juzhen SolidityStructTool automatically.
    *      Not to edit this code manually.
    */
    function update(TradeOperator storage _self, string _json) internal returns(bool succ) {
        if(LibJson.push(_json) == 0) {
            return false;
        }

        if (!_json.isJson()) {
            LibJson.pop();
            return false;
        }

        if (_json.jsonKeyExists("id"))
            _self.id = _json.jsonRead("id").toAddress();
        if (_json.jsonKeyExists("brokerId"))
            _self.brokerId = _json.jsonRead("brokerId").toUint();
        if (_json.jsonKeyExists("name"))
            _self.name = _json.jsonRead("name");
        if (_json.jsonKeyExists("department"))
            _self.department = _json.jsonRead("department");
        if (_json.jsonKeyExists("phone"))
            _self.phone = _json.jsonRead("phone");
        if (_json.jsonKeyExists("fax"))
            _self.fax = _json.jsonRead("fax");
        if (_json.jsonKeyExists("mobile"))
            _self.mobile = _json.jsonRead("mobile");
        if (_json.jsonKeyExists("email"))
            _self.email = _json.jsonRead("email");
        
        LibJson.pop();
        return true;
    }

    /**
    *@desc reset for TradeOperator
    *      Generated by juzhen SolidityStructTool automatically.
    *      Not to edit this code manually.
    */
    function reset(TradeOperator storage _self) internal {
        delete _self.id;
        delete _self.brokerId;
        delete _self.name;
        delete _self.department;
        delete _self.phone;
        delete _self.fax;
        delete _self.mobile;
        delete _self.email;
    }
}