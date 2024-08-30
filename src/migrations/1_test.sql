-- migrate:up
create table if not exists test (
    test_col varchar
)

-- migrate:down
drop table test