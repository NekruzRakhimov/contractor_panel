SELECT *
FROM contracts
WHERE id not in (select prev_contract_id from contracts)
  AND is_active = true;



SELECT id, status, requisites
FROM contracts;

SELECT id,
       type,
       status,
       requisites,
       manager,
       kam,
       supplier_company_manager,
       contract_parameters,
       with_temperature_conditions,
       products,
       discounts,
       comment,
       created_at,
       updated_at,
       is_individ,
       additional_agreement_number,
       ext_contract_code
FROM contracts;


SELECT id,
       type,
       status,
       requisites,
       manager,
       contract_parameters,
       created_at,
       updated_at,
       is_individ,
       additional_agreement_number,
       ext_contract_code
FROM contracts


SELECT id, contract_parameters
FROM contracts;


SELECT name
FROM contractors_contractor
WHERE id = 56;

SELECT id, requisites ->> 'beneficiary'
FROM contracts;


SELECT id,
       products,
       contract_parameters,
       requisites,
       manager,
       type,
       status,

      -- case when (requisites ->> 'beneficiary') IS NULL then '' else requisites ->> 'beneficiary' end AS suppler
       case when products IS NULL then '[]' else products end FROM contracts;



SELECT id, CAST(created_at AS text) FROM contracts;

SELECT id FROM co