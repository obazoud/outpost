interface Event {
  id: string;
  destination_id: string;
  topic: string;
  time: Date;
  successful_at: Date;
  metadata: any;
  data: any;
}

export type {
  Event,
};
