godbmig
=======

godbmig is an experimental projet for implementing migration for any languages / any projects which relates to any DB. I had two intentions for this project. 
a) create a tool which can track and generate migration files irrespective of any languages
b) integrate with any Go applications as a package to be used.


TODO
----

* Generate JSON.
* Create Schema_Migrations.
* Multi-DB support.
* Multi-DB Types.
* Make Transanctional whereever possible.
* Reading multiple actions in the Up / Down.
* Color Coded Output.
* Support YAML, XML storage too.
* Multi-support to read YAML / XML / JSON migrations.
* Generate SQL statements instead of executing directly.

* Read CONFIG either simple or from framework.

* MySQL - Rename Table.
* - sql: parameter.

Working
Generate
```
./godbmig add create_table test123 id:int name:string order:int status:bool
./godbmig add add_column test123 id:int
./godbmig add drop_column test123 id
./godbmig add sql 
```

Migration Up
```
./godbmig up
```

Not Working
Generate
```
./godbmig add rename_table old_test123 new_test123
./godbmig add add_index test123 id name
```

```
    # * add_column
    # * add_index
    # * add_timestamps
    # * create_table
    # * create_join_table
    # * remove_timestamps
    # * rename_column
    # * rename_index
    # * rename_table
```
godbmig down -1
```
