-- Migration: añadir campos de tenant (country, department, province, district, address, postal_code, logo_url)
ALTER TABLE tenant
ADD COLUMN IF NOT EXISTS country varchar(50) DEFAULT '';
ALTER TABLE tenant
ADD COLUMN IF NOT EXISTS department varchar(100) DEFAULT '';
ALTER TABLE tenant
ADD COLUMN IF NOT EXISTS province varchar(100) DEFAULT '';
ALTER TABLE tenant
ADD COLUMN IF NOT EXISTS district varchar(100) DEFAULT '';
ALTER TABLE tenant
ADD COLUMN IF NOT EXISTS address varchar(255) DEFAULT '';
ALTER TABLE tenant
ADD COLUMN IF NOT EXISTS postal_code varchar(20) DEFAULT '';
ALTER TABLE tenant
ADD COLUMN IF NOT EXISTS logo_url varchar(255) DEFAULT '';

-- Migration: añadir campos de respuesta en complaint
ALTER TABLE complaint
ADD COLUMN IF NOT EXISTS response_text text;
ALTER TABLE complaint
ADD COLUMN IF NOT EXISTS response_status varchar(20);
ALTER TABLE complaint
ADD COLUMN IF NOT EXISTS responder_id uuid;
ALTER TABLE complaint
ADD COLUMN IF NOT EXISTS response_sent_at timestamptz;

-- Index sobre responder_id si se desea
CREATE INDEX IF NOT EXISTS idx_complaint_responder_id ON complaint(responder_id);
