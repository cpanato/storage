-- Create index "document_types_proto_message_key" to table: "document_types"
CREATE UNIQUE INDEX `document_types_proto_message_key` ON `document_types` (`proto_message`);
-- Create index "nodes_proto_message_key" to table: "nodes"
CREATE UNIQUE INDEX `nodes_proto_message_key` ON `nodes` (`proto_message`);
-- Create index "persons_proto_message_key" to table: "persons"
CREATE UNIQUE INDEX `persons_proto_message_key` ON `persons` (`proto_message`);
-- Create index "tools_proto_message_key" to table: "tools"
CREATE UNIQUE INDEX `tools_proto_message_key` ON `tools` (`proto_message`);
-- Create index "properties_proto_message_key" to table: "properties"
CREATE UNIQUE INDEX `properties_proto_message_key` ON `properties` (`proto_message`);
-- Create index "source_data_proto_message_key" to table: "source_data"
CREATE UNIQUE INDEX `source_data_proto_message_key` ON `source_data` (`proto_message`);
-- Create index "edge_types_proto_message_key" to table: "edge_types"
CREATE UNIQUE INDEX `edge_types_proto_message_key` ON `edge_types` (`proto_message`);
-- Create index "external_references_proto_message_key" to table: "external_references"
CREATE UNIQUE INDEX `external_references_proto_message_key` ON `external_references` (`proto_message`);
-- Create index "metadata_proto_message_key" to table: "metadata"
CREATE UNIQUE INDEX `metadata_proto_message_key` ON `metadata` (`proto_message`);
-- Create index "node_lists_proto_message_key" to table: "node_lists"
CREATE UNIQUE INDEX `node_lists_proto_message_key` ON `node_lists` (`proto_message`);
