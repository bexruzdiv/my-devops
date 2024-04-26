do $$
    begin
        alter table sms_send drop column shipper_id;
    exception
        when others then
            RAISE NOTICE 'Already existed';
    end $$;