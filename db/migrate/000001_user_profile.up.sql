create table user_profiles (
	id uuid not null default gen_random_uuid(),
	created_at timestamptz not null default current_timestamp,
	created_by uuid null,
	updated_at timestamptz not null default current_timestamp,
	updated_by uuid null,
	deleted_at timestamptz null,
	deleted_by uuid null,
	profile_name varchar(500) not null,
	phone varchar(30) not null,
	email varchar(100) not null,
	constraint user_profiles_pk primary key(id)
);

create index user_profiles_email_idx on user_profiles using btree(phone);
create index user_profiles_email_phone_idx on user_profiles using btree(phone, email);
