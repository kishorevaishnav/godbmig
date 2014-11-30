godbmig
=======


TODO
----

* Generate JSON.
* Create Schema_Migrations.
* Multi-DB support.
* Multi-DB Types.
* Make Transanctional whereever possible.
* Reading multiple actions in the Up / Down.
* Color Coded Output

* Read CONFIG either simple or from framework.

* MySQL - Rename Table
* - sql: parameter
* 

godbmig add create_table test123 id:int name:string order:int status:bool
godbmig add rename_table old_test123 new_test123
godbmig add add_column test123 id:int
godbmig add drop_column test123 id
godbmig add sql 

godbmig up

godbmig down -1

