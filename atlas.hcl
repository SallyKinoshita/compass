schema "school_db" {
  charset = "utf8mb4"
  collate = "utf8mb4_0900_ai_ci"
}

table "facilitators" {
  schema = schema.school_db

  column "id" {
    type = int
    auto_increment = true
  }
  column "name" {
    type = varchar(255)
    null = true
  }
  primary_key {
    columns = [column.id]
  }
}

table "classrooms" {
  schema = schema.school_db

  column "id" {
    type = int
    auto_increment = true
  }
  column "name" {
    type = varchar(255)
    null = false
  }
  column "facilitator_id" {
    type = int
    null = false
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "facilitator_id" {
    columns = [column.facilitator_id]
    ref_columns = [table.facilitators.column.id]
    on_delete = "RESTRICT"
    on_update = "RESTRICT"
  }
}

table "students" {
  schema = schema.school_db

  column "id" {
    type = int
    auto_increment = true
  }
  column "name" {
    type = varchar(255)
    null = false
  }
  column "login_id" {
    type = varchar(255)
    null = false
  }
  column "classroom_id" {
    type = int
    null = false
  }
  primary_key {
    columns = [column.id]
  }
  index "index_login_id" {
    unique  = true
    columns = [column.login_id]
  }
  foreign_key "classroom_id" {
    columns = [column.classroom_id]
    ref_columns = [table.classrooms.column.id]
    on_delete = "RESTRICT"
    on_update = "RESTRICT"
  }
}
