type DestinationType = "rabbitmq" | "aws_sqs" | "webhook";

interface ConfigField {
  type: "text";
  label: string;
  description: string;
  key: string;
  required: boolean;
}

interface CredentialField extends ConfigField {
  sensitive?: boolean;
}

interface DestinationTypeReference {
  config_fields: ConfigField[];
  credential_fields: CredentialField[];
  instructions: string;
  label: string;
  description: string;
  icon: string;
  target: string;
}

interface Destination {
  id: string;
  type: DestinationType;
  config: Record<string, any>;
  credentials: Record<string, any>;
  label: string;
  description: string;
  icon: string;
  instructions: string;
  target: string;
  disabled_at: string;
}

export type {
  Destination,
  ConfigField,
  CredentialField,
  DestinationTypeReference,
};
