do $$
    begin
        ALTER TABLE sms_send ALTER COLUMN text TYPE VARCHAR(160);
    exception
        when others then
            RAISE NOTICE 'Already existed';
    end $$;