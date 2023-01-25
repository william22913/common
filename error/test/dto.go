package error_test

type TestDTO struct {
	Name string `json:"name" min:"5" max:"100" required:"insert,update" regex:"directory_name" reserved:"public,root,trash,backup,private"`
	Enum string `json:"enum" enum:"record_status" required:"insert"`
}
