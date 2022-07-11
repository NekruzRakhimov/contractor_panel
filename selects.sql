SELECT * FROM contracts WHERE id not in (select prev_contract_id from contracts) AND is_active = true;



SELECT  id, status, requisites FROM contracts ;

SELECT  id,type, status, requisites, manager, kam, supplier_company_manager, contract_parameters,with_temperature_conditions,
		products, discounts, comment,  created_at, updated_at,is_individ, additional_agreement_number,ext_contract_code FROM contracts;


SELECT  id, type, status, requisites, manager,  contract_parameters,
		 created_at, updated_at, is_individ, additional_agreement_number, ext_contract_code FROM contracts


SELECT  id, contract_parameters FROM contracts