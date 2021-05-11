variable "project" {
  type        = string
  default     = ""
  description = "Project, which could be your organization name or abbreviation, e.g. 'eg' or 'cp'"
}

variable "environment" {
  type        = string
  default     = ""
  description = "Environment, e.g. 'prod', 'staging', 'dev', 'pre-prod', 'UAT'"
}

variable "family" {
  type        = string
  default     = ""
  description = "Family, e.g. 'prod', 'staging', 'dev', OR 'source', 'build', 'test', 'deploy', 'release'"
}

variable "application" {
  type        = string
  default     = ""
  description = "Solution application, e.g. 'app' or 'jenkins'"
}

variable "model" {
  type        = string
  default     = ""
  description = "Additional model (e.g. `1`)"
}

variable "table_name" {
  type        = string
  default     = ""
  description = "overwrites autogenerated table name if set"
}

variable "enabled" {
  type        = bool
  default     = true
  description = "Set to false to prevent the module from creating any resources"
}

variable "enabled_global_secondary_index" {
  type        = bool
  default     = false
  description = "Set to false to prevent the module from creating any resources"
}

variable "autoscale_write_target" {
  type        = number
  default     = 75
  description = "The target value (in %) for DynamoDB write autoscaling"
}

variable "autoscale_read_target" {
  type        = number
  default     = 75
  description = "The target value (in %) for DynamoDB read autoscaling"
}

variable "autoscale_min_read_capacity" {
  type        = number
  default     = 1
  description = "DynamoDB autoscaling min read capacity"
}

variable "autoscale_max_read_capacity" {
  type        = number
  default     = 1000
  description = "DynamoDB autoscaling max read capacity"
}

variable "autoscale_min_write_capacity" {
  type        = number
  default     = 1
  description = "DynamoDB autoscaling min write capacity"
}

variable "autoscale_max_write_capacity" {
  type        = number
  default     = 1000
  description = "DynamoDB autoscaling max write capacity"
}

variable "autoscale_min_read_capacity_global_secondary_index" {
  type        = number
  default     = 1
  description = "DynamoDB autoscaling min read capacity"
}

variable "autoscale_max_read_capacity_global_secondary_index" {
  type        = number
  default     = 1000
  description = "DynamoDB autoscaling max read capacity"
}

variable "autoscale_min_write_capacity_global_secondary_index" {
  type        = number
  default     = 1
  description = "DynamoDB autoscaling min write capacity"
}

variable "autoscale_max_write_capacity_global_secondary_index" {
  type        = number
  default     = 1000
  description = "DynamoDB autoscaling max write capacity"
}

variable "billing_mode" {
  type        = string
  default     = "PROVISIONED"
  description = "DynamoDB Billing mode. Can be PROVISIONED or PAY_PER_REQUEST"
}

variable "enable_streams" {
  type        = bool
  default     = false
  description = "Enable DynamoDB streams"
}

variable "stream_view_type" {
  type        = string
  default     = ""
  description = "When an item in the table is modified, what information is written to the stream"
}

variable "enable_encryption" {
  type        = bool
  default     = false
  description = "Enable DynamoDB server-side encryption"
}

variable "enable_point_in_time_recovery" {
  type        = bool
  default     = false
  description = "Enable DynamoDB point in time recovery"
}

variable "hash_key" {
  type        = string
  description = "DynamoDB table Hash Key"
}

variable "hash_key_type" {
  type        = string
  default     = "S"
  description = "Hash Key type, which must be a scalar type: `S`, `N`, or `B` for (S)tring, (N)umber or (B)inary data"
}

variable "range_key" {
  type        = string
  default     = ""
  description = "DynamoDB table Range Key"
}

variable "range_key_type" {
  type        = string
  default     = "S"
  description = "Range Key type, which must be a scalar type: `S`, `N`, or `B` for (S)tring, (N)umber or (B)inary data"
}

variable "tags" {
  type        = "map"
  default     = {}
  description = "Additional tags for DynamoDB"
}

variable "ttl" {
  type        = string
  default     = "ttl"
  description = "DynamoDB table TTL attribute"
}

variable "enable_autoscaler" {
  type        = bool
  default     = true
  description = "Flag to enable/disable DynamoDB autoscaling"
}

variable "attributes" {
  type = list(object({
    name = string
    type = string
  }))
  default     = []
  description = "Additional DynamoDB attributes in the form of a list of mapped values"
}

variable "global_secondary_index" {
  type = list(object({
    hash_key           = string
    name               = string
    non_key_attributes = list(string)
    projection_type    = string
    range_key          = string
    read_capacity      = number
    write_capacity     = number
  }))
  default     = []
  description = "Additional global secondary indexes in the form of a list of mapped values"
}

variable "local_secondary_index" {
  type = list(object({
    name               = string
    non_key_attributes = list(string)
    projection_type    = string
    range_key          = string
  }))
  default     = []
  description = "Additional local secondary indexes in the form of a list of mapped values"
}

variable "regex_replace_chars" {
  type        = string
  default     = "/[^a-zA-Z0-9-]/"
  description = "Regex to replace chars with empty string in `namespace`, `environment`, `stage` and `name`. By default only hyphens, letters and digits are allowed, all other chars are removed"
}

variable "autoscaling_schedule_table_read_start" {
  type = list(object({
    cron         = string,
    min_capacity = number,
    max_capacity = number,
  }))

  default = []
}

variable "autoscaling_schedule_table_read_stop" {
  type = list(object({
    cron         = string,
    min_capacity = number,
    max_capacity = number,
  }))

  default = []
}

variable "autoscaling_schedule_table_write_start" {
  type = list(object({
    cron         = string,
    min_capacity = number,
    max_capacity = number,
  }))

  default = []
}

variable "autoscaling_schedule_table_write_stop" {
  type = list(object({
    cron         = string,
    min_capacity = number,
    max_capacity = number,
  }))

  default = []
}