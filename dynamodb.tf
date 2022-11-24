resource "aws_dynamodb_table" "employee" { 
   name = "employee"
   billing_mode = "PROVISIONED" 
   read_capacity = "5" 
   write_capacity = "2" 
   hash_key = "id"

   attribute { 
      name = "id" 
      type = "S" 
   } 

   ttl {
     enabled = true
     attribute_name = "ttl"
   }

  #  point_in_time_recovery { enabled = true } 
   # server_side_encryption { enabled = true } 
}