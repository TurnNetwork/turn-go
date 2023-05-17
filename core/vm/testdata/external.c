#define WASM_EXPORT __attribute__((visibility("default")))

#include <stddef.h>
#include <stdint.h>

uint8_t bub_gas_price(uint8_t gas_price[16]);
void bub_block_hash(int64_t num, uint8_t hash[32]);
uint64_t bub_block_number();
uint64_t bub_gas_limit();
uint64_t bub_gas();
int64_t bub_timestamp();
void bub_coinbase(uint8_t addr[20]);
uint8_t bub_balance(const uint8_t addr[20], uint8_t balance[16]);
void bub_origin(uint8_t addr[20]);
void bub_caller(uint8_t addr[20]);
uint8_t bub_call_value(uint8_t val[16]);
void bub_address(uint8_t addr[20]);
void bub_sha3(const uint8_t *src, size_t srcLen, uint8_t *dest, size_t destLen);
uint64_t bub_caller_nonce();
int32_t bub_transfer(const uint8_t to[20], const uint8_t *amount, size_t len);
void bub_set_state(const uint8_t *key, size_t klen, const uint8_t *value, size_t vlen);
size_t bub_get_state_length(const uint8_t *key, size_t klen);
int32_t bub_get_state(const uint8_t *key, size_t klen, uint8_t *value, size_t vlen);
size_t bub_get_input_length();
void bub_get_input(const uint8_t *value);
size_t bub_get_call_output_length();
void bub_get_call_output(const uint8_t *value);
void bub_return(const uint8_t *value, const size_t len);
void bub_revert();
void bub_panic();
void bub_debug(uint8_t *dst, size_t len);
int32_t bub_call(const uint8_t to[20], const uint8_t *args, size_t args_len, const uint8_t *value, size_t value_len, const uint8_t *call_cost, size_t call_cost_len);
int32_t bub_delegate_call(const uint8_t to[20], const uint8_t *args, size_t args_len, const uint8_t *call_cost, size_t call_cost_len);
//int32_t bub_static_call(const uint8_t to[20], const uint8_t* args, size_t argsLen, const uint8_t* callCost, size_t callCostLen);
int32_t bub_destroy(const uint8_t to[20]);
int32_t bub_migrate(uint8_t new_addr[20], const uint8_t *args, size_t args_len, const uint8_t *value, size_t value_len, const uint8_t *call_cost, size_t call_cost_len);
int32_t bub_clone_migrate(const uint8_t old_addr[20], uint8_t new_addr[20], const uint8_t *args, size_t args_len, const uint8_t *value, size_t value_len, const uint8_t *call_cost, size_t call_cost_len);
void bub_event(const uint8_t *topic, size_t topic_len, const uint8_t *args, size_t args_len);
int32_t bub_ecrecover(const uint8_t hash[32], const uint8_t* sig, const uint8_t sig_len, uint8_t addr[20]);
void bub_ripemd160(const uint8_t *input, uint32_t input_len, uint8_t addr[20]);
void bub_sha256(const uint8_t *input, uint32_t input_len, uint8_t hash[32]);

// u128
size_t rlp_u128_size(uint64_t heigh, uint64_t low);
void bub_rlp_u128(uint64_t heigh, uint64_t low, void * dest);

// bytes
size_t rlp_bytes_size(const void *data, size_t len);
void bub_rlp_bytes(const void *data, size_t len, void * dest);

// list
size_t rlp_list_size(size_t len);
void bub_rlp_list(const void *data, size_t len, void * dest);

// get code length
size_t bub_contract_code_length(const uint8_t addr[20]);

// get code
int32_t bub_contract_code(const uint8_t addr[20], uint8_t *code,
                             size_t code_length);

// deploy new contract
int32_t bub_deploy(uint8_t new_addr[20], const uint8_t *args,
                      size_t args_len, const uint8_t *value, size_t value_len,
                      const uint8_t *call_cost, size_t call_cost_len);

// clone new contract
int32_t bub_clone(const uint8_t old_addr[20], uint8_t new_addr[20],
                     const uint8_t *args, size_t args_len, const uint8_t *value,
                     size_t value_len, const uint8_t *call_cost,
                     size_t call_cost_len);

uint8_t global_info[10] = {};

size_t rlp_unsigned(uint32_t data){
  int valid = 0, i = 1;
  for(int j = 24; j >= 0; j -= 8){
    uint8_t one = data >> j;
    if(one && !valid) valid = 1;
    if(valid) {
      global_info[i] = one;
      i++;
    }
  }
  global_info[0] = 0x80 + i - 1;
  return i;
}

WASM_EXPORT
void bub_gas_price_test() {
    uint8_t gas[32] = {0};
    uint8_t len = bub_gas_price(gas);
    bub_return(gas, len);
}

WASM_EXPORT
void bub_block_hash_test() {
  uint8_t hash[32];
  bub_block_hash(0, hash);
  bub_return(hash, sizeof(hash));
}

WASM_EXPORT
void bub_block_number_test() {
  uint64_t num = bub_block_number();
  bub_return((uint8_t*)&num, sizeof(num));
}
WASM_EXPORT
void bub_gas_limit_test() {
  uint64_t num = bub_gas_limit();
  bub_return((uint8_t*)&num, sizeof(num));
}

WASM_EXPORT
void bub_gas_test() {
  uint64_t num = bub_gas();
  bub_return((uint8_t*)&num, sizeof(num));
}

WASM_EXPORT
void bub_timestamp_test() {
  uint64_t num = bub_timestamp();
  bub_return((uint8_t*)&num, sizeof(num));
}

WASM_EXPORT
void bub_coinbase_test() {
  uint8_t hash[20];
  bub_coinbase(hash);
  bub_return(hash, sizeof(hash));
}

WASM_EXPORT
void bub_balance_test() {
  uint8_t hash[32] = {1};
  uint8_t balance[32] = {0};
  uint8_t len = bub_balance(hash, balance);
  bub_return(balance, len);
}

WASM_EXPORT
void bub_origin_test() {
  uint8_t hash[20];
  bub_origin(hash);
  bub_return(hash, sizeof(hash));
}

WASM_EXPORT
void bub_caller_test() {
  uint8_t hash[20];
  bub_caller(hash);
  bub_return(hash, sizeof(hash));
}

WASM_EXPORT
void bub_call_value_test() {
  uint8_t hash[32];
  uint8_t len = bub_call_value(hash);
  bub_return(hash, len);
}

WASM_EXPORT
void bub_address_test() {
  uint8_t hash[20];
  bub_address(hash);
  bub_return(hash, sizeof(hash));
}

WASM_EXPORT
void bub_sha3_test() {
  uint8_t data[1024];
  size_t len = bub_get_input_length();
  bub_get_input(data);
  uint8_t hash[32];
  bub_sha3(data, len, hash, 32);
  bub_return(hash, sizeof(hash));
}


WASM_EXPORT
void bub_caller_nonce_test() {
  uint64_t num = bub_caller_nonce();
  bub_return((uint8_t*)&num, sizeof(num));
}

WASM_EXPORT
void bub_set_state_test() {
  uint8_t data[1024];
  size_t len = bub_get_input_length();
  bub_get_input(data);
  bub_set_state((uint8_t*)"key", 3, data, len);
}

WASM_EXPORT
void bub_get_state_test() {
  uint8_t data[1024];
  size_t len = bub_get_state_length((uint8_t*)"key", 3);
  bub_get_state((uint8_t*)"key", 3, data, 1024);
  bub_return(data, len);
}

WASM_EXPORT
void bub_get_call_output_test() {
  uint8_t data[1024];
  size_t len = bub_get_call_output_length();
  bub_get_call_output(data);
  bub_return(data, len);
}

WASM_EXPORT
void bub_revert_test() {
  bub_revert();
}

WASM_EXPORT
void bub_panic_test() {
  bub_panic();
}

WASM_EXPORT
void bub_debug_test() {
  uint8_t data[1024];
  size_t len = bub_get_input_length();
  bub_get_input(data);
  bub_debug(data, len);
}

WASM_EXPORT
void bub_transfer_test() {
  uint8_t data[1024];
  size_t len = bub_get_input_length();
  bub_get_input(data);
  uint8_t value = 1;
  bub_transfer(data, &value, 1);
  bub_return(&value, 1);
}

WASM_EXPORT
void bub_call_contract_test() {
  uint8_t addr[20] = {1, 2, 4}; // don't change it
  uint8_t data[1024];
  size_t datalen = bub_get_input_length();
  bub_get_input(data);
  uint8_t gas = 100000;
  uint8_t value = 2;
  bub_call(addr, data, datalen, &value, 1, &gas, 5);
}

WASM_EXPORT
void bub_delegate_call_contract_test () {
    uint8_t addr[20] = {1, 2, 4}; // don't change it
    uint8_t data[1024];
    size_t datalen = bub_get_input_length();
    bub_get_input(data);
    uint8_t gas = 100000;
    bub_delegate_call(addr, data, datalen, &gas, 5);

}

//WASM_EXPORT
//void bub_static_call_contract_test () {
//   uint8_t addr[20] = {1, 2, 4}; // don't change it
//   uint8_t data[1024];
//   size_t datalen = bub_get_input_length();
//   bub_get_input(data);
//   uint8_t gas = 100000;
//   bub_static_call(addr, &data, datalen, &gas, 5);
//}

WASM_EXPORT
void bub_destroy_contract_test () {
    uint8_t addr[20] = {1, 2, 6};
    bub_destroy(addr);
}

WASM_EXPORT
void bub_migrate_contract_test () {
    uint8_t newAddr[20];
    uint8_t data[1024];
    size_t datalen = bub_get_input_length();
    bub_get_input(data);
    uint32_t gas = 1000000;
    size_t rlp_len = rlp_unsigned(gas);
    uint8_t value = 2;
    bub_migrate(newAddr, data, datalen, &value, 1, &global_info[1], rlp_len -1);
    bub_return(newAddr, 20);
}

WASM_EXPORT
void bub_clone_migrate_contract_test() {
    // get input
    uint8_t data[100];
    size_t datalen = bub_get_input_length();
    bub_get_input(data);

    uint32_t gas = 1000000;
    size_t rlp_len = rlp_unsigned(gas);
    uint8_t value = 2;

    uint8_t newAddr[20];
    uint8_t oldAddr[20] = {1, 2, 3};
    bub_clone_migrate(oldAddr, newAddr, data, datalen, &value, 1, &global_info[1], rlp_len -1);
    bub_return(newAddr, 20);
}

WASM_EXPORT
void bub_clone_migrate_contract_error_test() {
    // get input
    uint8_t data[100];
    size_t datalen = bub_get_input_length();
    bub_get_input(data);

    uint32_t gas = 1000000;
    size_t rlp_len = rlp_unsigned(gas);
    uint8_t value = 2;

    uint8_t newAddr[20];
    uint8_t oldAddr[20] = {1, 2, 3};
    bub_clone_migrate(oldAddr, newAddr, data, datalen, &value, 1, &global_info[1], rlp_len -1);
    bub_return(newAddr, 20);
}

WASM_EXPORT
void bub_event0_test () {

    uint8_t data[1024];
    size_t len = bub_get_input_length();
    bub_get_input(data);

    // empty topic
    uint8_t topics[1] = {0};

    bub_event(topics, 0, data, len);
}

WASM_EXPORT
void bub_event3_test () {

    uint8_t data[1024];
    size_t len = bub_get_input_length();
    bub_get_input(data);

    // rlp([topic1, topic2, topic3])
    uint8_t topics[10] = {201, 130, 116, 49, 130, 116, 50, 130, 116, 51};

    bub_event(topics, 10, data, len);
}
void bub_sha256(const uint8_t *input, uint32_t input_len, uint8_t hash[32]);
void bub_ripemd160(const uint8_t *input, uint32_t input_len, uint8_t addr[20]);
int32_t bub_ecrecover(const uint8_t hash[32], const uint8_t* sig, const uint8_t sig_len, uint8_t addr[20]);
WASM_EXPORT
void bub_sha256_test() {
    uint8_t input[3] = {1,2,3};
//    uint8_t hash[32] = {3,144,88,198,242,192,203,73,44,83,59,10,77,20,239,119,204,15,120,171,204,206,213,40,125,132,161,162,1,28,251,129};
    uint8_t res[32] = {0};
    bub_sha256(input, 3, res);
    bub_return(res, 32);
}

WASM_EXPORT
void bub_ripemd160_test() {
    uint8_t input[3] = {1,2,3};
//    uint8_t addr[20] = {121,249,1,218,38,9,240,32,173,173,191,46,95,104,161,108,140,63,125,87};
    uint8_t res[20] = {0};
    bub_ripemd160(input, 3, res);
    bub_return(res, 20);
}

WASM_EXPORT
void bub_ecrecover_test() {
    uint8_t hash[32] = {65,177,160,100,151,82,175,27,40,179,220,41,161,85,110,238,120,30,74,76,58,31,127,83,249,15,168,52,222,9,140,77};
    uint8_t sig[65] = {209,85,233,67,5,175,126,7,221,140,50,135,62,92,3,203,149,201,224,89,96,239,133,190,156,7,246,113,218,88,199,55,24,193,154,220,57,122,33,26,169,232,126,81,158,32,56,197,163,182,88,97,141,179,53,247,79,128,11,142,12,254,239,68,1};
//    uint8_t addr[20] = {151,14,129,40,171,131,78,142,172,23,171,142,56,18,240,16,103,140,247,145};
    uint8_t res[20] = {0};
    bub_ecrecover(hash, sig, 65, res);
    bub_return(res, 20);
}

WASM_EXPORT
void rlp_u128_size_test(){
  uint64_t heigh = 0x0123456789abcdefULL;
  uint64_t low = 0xfedcba9876543210ULL;
  size_t append_length = rlp_u128_size(heigh, low);
  uint8_t res[8] = {0};
  for(int i = 0; i < 8; i++){
    res[i] = append_length >> (i * 8);
  }
  bub_return(res, 8);
}

WASM_EXPORT
void bub_rlp_u128_test(){
  uint64_t heigh = 0x0123456789abcdefULL;
  uint64_t low = 0xfedcba9876543210ULL;
  uint8_t res[17] = {0};
  bub_rlp_u128(heigh, low, res);
  bub_return(res, 17);
}

WASM_EXPORT
void rlp_bytes_size_test(){
  uint8_t data[16] = {0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x00a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f};
  size_t append_length = rlp_bytes_size(data, 16);
  uint8_t res[8] = {0};
  for(int i = 0; i < 8; i++){
    res[i] = append_length >> (i * 8);
  }
  bub_return(res, 8);
}

WASM_EXPORT
void bub_rlp_bytes_test(){
  uint8_t data[16] = {0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x00a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f};
  uint8_t res[17] = {0};
  bub_rlp_bytes(data, 16, res);
  bub_return(res, 17);
}

WASM_EXPORT
void rlp_list_size_test(){
  uint8_t data[16] = {0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x00a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f};
  size_t append_length = rlp_list_size(sizeof(data));
  uint8_t res[8] = {0};
  for(int i = 0; i < 8; i++){
    res[i] = append_length >> (i * 8);
  }
  bub_return(res, 8);
}

WASM_EXPORT
void bub_rlp_list_test(){
  uint8_t data[16] = {0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x00a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f};
  uint8_t res[17] = {0};
  bub_rlp_list(data, 16, res);
  bub_return(res, 17);
}

WASM_EXPORT
void bub_contract_code_length_test(){
  uint8_t contractAddr[20] = {1, 2, 3};
  size_t length = bub_contract_code_length(contractAddr);
  size_t rlp_len = rlp_unsigned(length);
  bub_return(global_info, rlp_len);
}

WASM_EXPORT
void bub_contract_code_test(){
  uint8_t contractAddr[20] = {1, 2, 3};
  size_t length = bub_contract_code_length(contractAddr);
  uint8_t code[16] = {};
  bub_contract_code(contractAddr, code, 16);
  bub_return(code, 16);
}

WASM_EXPORT
void bub_deploy_test () {
    uint8_t data[1024];
    size_t datalen = bub_get_input_length();
    bub_get_input(data);

    uint32_t gas = 1000000;
    size_t rlp_len = rlp_unsigned(gas);
    uint8_t value = 2;
    uint8_t newAddr[20] = {};
    bub_deploy(newAddr, data, datalen, &value, 1, &global_info[1], rlp_len - 1);
    bub_return(newAddr, 20);
}

WASM_EXPORT
void bub_clone_test() {
    // get input
    uint8_t data[100];
    size_t datalen = bub_get_input_length();
    bub_get_input(data);

    uint32_t gas = 1000000;
    size_t rlp_len = rlp_unsigned(gas);
    uint8_t value = 2;

    uint8_t newAddr[20] = {};
    uint8_t oldAddr[20] = {1, 2, 3};
    bub_clone(oldAddr, newAddr, data, datalen, &value, 1, &global_info[1], rlp_len - 1);
    bub_return(newAddr, 20);
}

WASM_EXPORT
void bub_clone_error_test() {
    // get input
    uint8_t data[100];
    size_t datalen = bub_get_input_length();
    bub_get_input(data);

    uint32_t gas = 1000000;
    size_t rlp_len = rlp_unsigned(gas);
    uint8_t value = 2;

    uint8_t newAddr[20] = {};
    uint8_t oldAddr[20] = {1, 2, 3};
    bub_clone(oldAddr, newAddr, data, datalen, &value, 1, &global_info[1], rlp_len - 1);
    bub_return(newAddr, 20);
}