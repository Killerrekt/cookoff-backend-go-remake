Create table if not exists users (
    name text,
    email text primary key,
    regNo text,
    refreshToken text,
    userRole text default 'user',
    isActive boolean default false,
    isRoundActive boolean default false,
    roundQualified integer default 0,
    password text,
    tokenVersion integer default 0,
    score integer default 0,
    submissionTime timestamp);