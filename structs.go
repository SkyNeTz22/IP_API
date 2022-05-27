package main

type Config struct{
	MySql	MySql	`json:"mySql"`
}

type MySql struct{
	Username	string	`json:"username"`
	Password	string	`json:"password"`
	Network		string	`json:"net"`
	HostName	string	`json:"hostName"`
	DbName		string	`json:"dbName"`
}