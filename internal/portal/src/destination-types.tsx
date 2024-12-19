import useSWR from "swr";
import { DestinationTypeReference } from "./typings/Destination";

const targets = {
  webhook: "url",
  rabbitmq: "exchange",
  aws: "queue_url",
};

export function useDestinationTypes(): Record<
  string,
  DestinationTypeReference
> {
  const { data } =
    useSWR<Record<string, DestinationTypeReference>>("destination-types");
  if (!data) {
    return {};
  }
  return Object.values(data).reduce((acc, type) => {
    acc[type.type] = {
      ...type,
      target: targets[type.type as keyof typeof targets],
    };
    return acc;
  }, {} as Record<string, DestinationTypeReference>);
}

export function useDestinationType(
  type: string | undefined
): DestinationTypeReference | undefined {
  const destination_types = useDestinationTypes();

  if (!type) {
    return undefined;
  }

  return destination_types[type];
}
