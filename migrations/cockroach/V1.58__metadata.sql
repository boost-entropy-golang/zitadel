CREATE TABLE auth.metadata (
   aggregate_id TEXT,

   key TEXT,
   value BYTES,

   resource_owner TEXT,
   creation_date TIMESTAMPTZ,
   change_date TIMESTAMPTZ,
   sequence BIGINT,

   PRIMARY KEY (aggregate_id, key)
);

CREATE TABLE management.metadata (
    aggregate_id TEXT,

    key TEXT,
    value BYTES,

    resource_owner TEXT,
    creation_date TIMESTAMPTZ,
    change_date TIMESTAMPTZ,
    sequence BIGINT,

    PRIMARY KEY (aggregate_id, key)
);
