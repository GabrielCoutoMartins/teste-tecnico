create table IF NOT EXISTS tasks (
	id serial primary key,
	titulo varchar,
	descricao varchar,
	status varchar,
	criadoEm date default now()
)