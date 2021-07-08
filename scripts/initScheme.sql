drop table if exists Users cascade;
drop table if exists Groups cascade;
drop table if exists Students cascade;
drop table if exists Professors cascade;
drop table if exists Competenties cascade;

drop type if exists user_roles;

create type users_roles as enum ('admin', 'methodist', 'student', 'professor', 'curator', 'guest');

create unlogged table Users
(
	user_id 		uuid 	primary key not null,
	role 			users_roles not null,
	email 			citext,
	password_hash 	bytea,
	name 			varchar(128) not null,
	surname 		varchar(128) not null,
	patronymic 		varchar(128) not null,
	phone 			varchar(18),
	birth_date 		date,
	about 			text
)

create unlogged table Groups
(
	group_nmb 			int2 primary key not null,
	group_elder_id 		uuid foreign key references users(user_id), /*Староста*/
	timetable_id 		uuid, 										/*Расписание*/
	group_curator_id 	uuid foreign key references users(user_id), /*Куратор*/
	semester 			int2,
	students_cnt 		int2
)

create unlogged table Students
(
	user_id 		uuid foreign key references users(user_id) not null,
	org_curator_id 	uuid foreign key references users(user_id) not null,
	group_nmb 		int2 foreign key references groups(group_nmb) not null,
	admission_date 	date not null,
	is_graduated 	bool,
	in_academ 		bool
)

create unlogged table Professors
(
	user_id 		uuid foreign key references users(user_id) not null,
	seniority 		int2, 			/*Стаж, мес.*/
	academic_degree varchar(256), 	/*Учёная степень*/
	rank 			varchar(32), 	/*Звание*/
	contest_date 	date, 			/*Дата конкурса*/
	is_combining 	bool, 			/*Совмещение*/
	shared_hours 	int2, 			/*Общие часы*/
	work_rate 		int4, 			/*Ставка в руб.*/
	work_time 		varchar(32), 	/*Время работы на кафедре*/
	disciplines 	varchar(128)[]
)

create unlogged table Competenties
(
	competention varchar(64) primary key not null,
	users_ids uuid[]
)
