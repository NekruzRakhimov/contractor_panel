SELECT * FROM contracts WHERE id not in (select prev_contract_id from contracts) AND is_active = true;