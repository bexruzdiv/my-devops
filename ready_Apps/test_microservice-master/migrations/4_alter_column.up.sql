do $$
    begin
        ALTER TABLE sms_send ALTER COLUMN text TYPE TEXT;
    exception
        when others then
            RAISE NOTICE 'Already existed';
    end $$;