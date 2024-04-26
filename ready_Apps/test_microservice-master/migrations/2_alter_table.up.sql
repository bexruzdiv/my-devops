do $$
    begin
        alter table sms_send add column shipper_id uuid;
        alter table sms_send add column send_count smallint default 0;
    exception
        when others then
            RAISE NOTICE 'Already existed';
    end $$;