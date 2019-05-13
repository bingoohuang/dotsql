-- name: FindUser
select username, password, fullname, email, mobile, detail from typhon_admin_user where username = ?;

-- name: AddUser
insert into typhon_admin_user(username, password, fullname, email, mobile, detail) values(?,?,?,?,?,?);


-- name: ClientLogSql
insert into typhon_cur_client(app_id, conf_file, crc, ip, sync_time)
values (?, ?, ?, ?, ?)
on duplicate key update crc = ?, sync_time= ?;

-- name: ClientLogSql-postgres
insert into typhon_cur_client(app_id, conf_file, crc, ip, sync_time)
values (?, ?, ?, ?, ?)
on conflict (app_id, conf_file, ip) do update set crc = ?, sync_time = ?;



-- name: createTeams
CREATE TABLE teams (id INTEGER PRIMARY KEY ASC, name TEXT);

-- name: addTeam
INSERT INTO teams (name) VALUES (?);


-- name: findTeam
SELECT ID, NAME FROM teams WHERE id = ?;

-- name: selectTeams
SELECT ID, NAME FROM teams;


