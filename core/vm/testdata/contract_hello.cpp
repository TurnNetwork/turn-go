#include <bubble/bubble.hpp>
#include <string>
using namespace bubble;


class message {
   public:
      message(){}
      message(const std::string &p_head):head(p_head){}
   private:
      std::string head;
      BUBBLE_SERIALIZE(message, (head))
};

class my_message : public message {
   public:
      my_message(){}
      my_message(const std::string &p_head, const std::string &p_body, const std::string &p_end):message(p_head), body(p_body), end(p_end){}
   private:
      std::string body;
      std::string end;
      BUBBLE_SERIALIZE_DERIVED(my_message, message, (body)(end))
};

CONTRACT hello : public bubble::Contract{
   public:
     

     ACTION void init(){}
      
      ACTION std::vector<my_message> add_message(const my_message &one_message){
          arr.self().push_back(one_message);
          return arr.self();
      }
      CONST std::vector<my_message> get_message(const std::string &name){
          return arr.self();
      }

      CONST uint64_t get_vector_size(){
          DEBUG("hello get_vector_size");
          return arr.self().size();
      }


   private:
      bubble::StorageType<"arr"_n, std::vector<my_message>> arr;
};

BUBBLE_DISPATCH(hello, (init)(add_message)(get_message)(get_vector_size))
