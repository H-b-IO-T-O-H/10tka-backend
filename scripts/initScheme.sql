create extension if not exists "uuid-ossp";
create extension if not exists citext;

drop table if exists Users cascade;
drop table if exists Groups cascade;
drop table if exists Students cascade;
drop table if exists Professors cascade;
drop table if exists Competenties cascade;
drop table if exists DisciplinesMaterials cascade;
drop table if exists Disciplines cascade;
drop table if exists Organizations cascade;
drop table if exists Audiences cascade;
drop table if exists AudiencesMaterials cascade;
drop table if exists Timetables cascade;
drop table if exists Weeks cascade;
drop table if exists Days cascade;
drop table if exists Lessons cascade;
drop table if exists Posts cascade;
drop table if exists PostsComments cascade;
drop table if exists PostsLikes cascade;

drop type if exists users_roles;
drop type if exists disp_types;
drop type if exists week_types;
drop type if exists lessons_types;
drop type if exists tag_types;

create type users_roles as enum ('admin', 'methodist', 'student', 'professor', 'curator', 'guest');
create type disp_types as enum ('exam', 'offset', 'diff_offset');
create type week_types as enum ('numerator', 'denominator');
create type lessons_types as enum ('seminar', 'lecture', 'lab', 'homework', 'rcontrol', 'consultation', 'exam', 'free');
create type tag_types as enum ('important', 'general', 'education');

create unlogged table Users
(
    user_id       serial
        primary key            not null,
    role          users_roles  not null,
    email         citext collate "C",
    password_hash bytea,
    name          varchar(128) not null,
    surname       varchar(128) not null,
    patronymic    varchar(128) not null,
    phone         varchar(18),
    birth_date    date,
    about         text,
    avatar_path   varchar(128) default null
);

create unlogged table Groups
(
    group_nmb        int2 primary key not null,
    group_elder_id   serial
        references Users (user_id)    not null, /*Староста*/
    timetable_id     uuid             not null, /*Расписание*/
    group_curator_id serial
        references Users (user_id)    not null, /*Куратор*/
    semester         int2,
    students_cnt     int2
);

create unlogged table Students
(
    user_id        serial
        references Users (user_id)    not null,
    org_curator_id serial
        references Users (user_id)    not null,
    group_nmb      int2
        references Groups (group_nmb) not null,
    admission_date date               not null,
    is_graduated   bool,
    in_academ      bool
);

create unlogged table Professors
(
    user_id         serial
        references Users (user_id) not null,
    seniority       int2           not null, /*Стаж, мес.*/
    academic_degree varchar(256), /*Учёная степень*/
    prof_rank       varchar(32), /*Звание*/
    contest_date    date, /*Дата конкурса*/
    is_combining    bool, /*Совмещение*/
    shared_hours    int2, /*Общие часы*/
    work_rate       int4, /*Ставка в руб.*/
    work_time       varchar(32), /*Время работы на кафедре*/
    disciplines     varchar(128)[]
);

create unlogged table Competenties
(
    competence varchar(64) primary key not null,
    users_ids  int4[]
);

create unlogged table Disciplines
(
    disc_name    varchar(128) primary key not null,
    semester     int2                     not null,
    is_part_time bool,
    is_secret    bool,
    disp_type    disp_types               not null,
    competencies varchar(64)[]
);

create unlogged table DisciplinesMaterials
(
    mat_name        varchar(64) primary key not null,
    disc_name       varchar(128)
        references Disciplines (disc_name)  not null,
    mat_description text,
    mat_filename    varchar(128),
    mat_cnt         int2
);

create unlogged table Organizations
(
    org_name       varchar(64) primary key not null,
    org_curator_id serial
        references Users (user_id)         not null,
    department     varchar(16)             not null
);

create unlogged table Audiences
(
    aud_name            varchar(16) primary key not null,
    responsible_user_id serial
        references Users (user_id)              not null,
    is_secret           bool,
    capacity            int2,
    aud_employment      varchar(128),
    about               text
);

create unlogged table AudiencesMaterials
(
    mat_name        varchar(64) primary key not null,
    aud_name        varchar(16)
        references Audiences (aud_name)     not null,
    responsible_id  serial
        references Users (user_id)          not null,
    mat_description text,
    mat_serial_nmb  varchar(128)
);

create unlogged table Timetables
(
    group_nmb  int2
        references Groups (group_nmb) not null,
    semester   int2                   not null,
    weeks_nmbs int2[18]
);

create unlogged table Lessons
(
    lesson_code varchar(6) primary key     not null,
    disc_name   varchar(128)
        references Disciplines (disc_name) not null,
    aud_name    varchar(16)
        references Audiences (aud_name)    not null,
    lesson_type lessons_types              not null
);

create unlogged table Days
(
    day_code varchar(4) primary key                      not null,
    l1_code  varchar(6) references Lessons (lesson_code) not null,
    l2_code  varchar(6) references Lessons (lesson_code) not null,
    l3_code  varchar(6) references Lessons (lesson_code) not null,
    l4_code  varchar(6) references Lessons (lesson_code) not null,
    l5_code  varchar(6) references Lessons (lesson_code) not null,
    l6_code  varchar(6) references Lessons (lesson_code) not null,
    l7_code  varchar(6) references Lessons (lesson_code) not null,
    l8_code  varchar(6) references Lessons (lesson_code) not null
);

create unlogged table Weeks
(
    week_nmb  int2 primary key                      not null,
    week_type week_types                            not null,
    monday    varchar(4) references Days (day_code) not null,
    tuesday   varchar(4) references Days (day_code) not null,
    wednesday varchar(4) references Days (day_code) not null,
    thursday  varchar(4) references Days (day_code) not null,
    friday    varchar(4) references Days (day_code) not null,
    saturday  varchar(4) references Days (day_code) not null,
    sunday    varchar(4) references Days (day_code) not null
);

create unlogged table Posts
(
    id        serial primary key,
    author_id serial
        references Users (user_id) not null,
    title     varchar(256)         not null,
    tag_type  tag_types default 'general',
    content   text                 not null,
    is_edited bool      default false,
    created   timestamp            not null,
    comments  bool      default false
);

create unlogged table PostsComments
(
    id        serial primary key,
    post_id   serial
        references Posts (id)      not null,
    author_id serial
        references Users (user_id) not null,
    parent    int   default 0,
    message   text                 not null,
    is_edited bool  default false,
    created   timestamp            not null,
    path      int[] default array []::int[]
);

create unlogged table PostsLikes
(
    post_id serial references Posts (id),
    user_id serial references Users (user_id)
)