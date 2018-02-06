-- Copyright 2017 The LUCI Authors.
--
-- Licensed under the Apache License, Version 2.0 (the "License");
-- you may not use this file except in compliance with the License.
-- You may obtain a copy of the License at
--
--      http://www.apache.org/licenses/LICENSE-2.0
--
-- Unless required by applicable law or agreed to in writing, software
-- distributed under the License is distributed on an "AS IS" BASIS,
-- WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
-- See the License for the specific language governing permissions and
-- limitations under the License.

-- Required fields are not enforced by this schema.
-- The Machine Database will enforce any such constraints.

ALTER TABLE vms DROP COLUMN state;
ALTER TABLE physical_hosts DROP COLUMN state;
ALTER TABLE machines DROP COLUMN state;
ALTER TABLE vlans DROP COLUMN state;
ALTER TABLE platforms DROP COLUMN state;
ALTER TABLE oses DROP COLUMN state;
ALTER TABLE switches DROP COLUMN state;
ALTER TABLE racks DROP COLUMN state;
ALTER TABLE datacenters DROP COLUMN state;