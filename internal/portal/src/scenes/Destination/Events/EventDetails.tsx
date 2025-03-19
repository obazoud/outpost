import { Route, Routes, useLocation, useParams } from "react-router-dom";
import Button from "../../../common/Button/Button";
import { CloseIcon, PreviousIcon } from "../../../common/Icons";
import useSWR from "swr";
import { Event, Delivery } from "../../../typings/Event";

const EventDetails = ({
  navigateEvent,
}: {
  navigateEvent: (path: string, params?: any) => void;
}) => {
  return (
    <Routes>
      <Route path="/" element={<EventData navigateEvent={navigateEvent} />} />
      <Route
        path="/attempts"
        element={<EventAttempts navigateEvent={navigateEvent} />}
      />
      <Route
        path="/attempts/:delivery_id"
        element={<EventAttemptDetails navigateEvent={navigateEvent} />}
      />
    </Routes>
  );
};

export default EventDetails;

const EventData = ({
  navigateEvent,
}: {
  navigateEvent: (path: string, params?: any) => void;
}) => {
  const { destination_id: destinationId, event_id: eventId } = useParams();

  const { data: event } = useSWR<Event>(
    `destinations/${destinationId}/events/${eventId}`
  );

  return (
    <div className="drawer">
      <div className="drawer__header">
        <div className="drawer__header-tabs">
          <Button onClick={() => navigateEvent(`/${eventId}`)}>Event</Button>
          <Button onClick={() => navigateEvent(`/${eventId}/attempts`)}>
            Attempts
          </Button>
        </div>

        <div>
          <Button minimal onClick={() => navigateEvent("/")}>
            <CloseIcon />
          </Button>
        </div>
      </div>

      <div className="drawer__body">
        <div className="event-data">
          <pre>{JSON.stringify(event, null, 2)}</pre>
        </div>
      </div>
    </div>
  );
};

const EventAttempts = ({
  navigateEvent,
}: {
  navigateEvent: (path: string, params?: any) => void;
}) => {
  const { event_id: eventId } = useParams();

  const { data: deliveries } = useSWR<Delivery[]>(
    `events/${eventId}/deliveries`
  );

  return (
    <div className="drawer">
      <div className="drawer__header">
        <div className="drawer__header-tabs">
          <Button onClick={() => navigateEvent(`/${eventId}`)}>Event</Button>
          <Button onClick={() => navigateEvent(`/${eventId}/attempts`)}>
            Attempts
          </Button>
        </div>

        <div>
          <Button minimal onClick={() => navigateEvent("/")}>
            <CloseIcon />
          </Button>
        </div>
      </div>

      <div className="drawer__body">
        <div className="event-attempts">
          <ol className="event-attempts__list">
            {deliveries?.map((delivery, index) => {
              const deliveryIndex = deliveries.length - index;
              return (
                <li
                  key={String(delivery.delivered_at)}
                  className="event-attempts__list-item"
                >
                  <button
                    type="button"
                    onClick={() => {
                      navigateEvent(`/${eventId}/attempts/${delivery.id}`, {
                        delivery,
                        deliveryIndex,
                      });
                    }}
                  >
                    <span className="visually-hidden">Go to attempt</span>
                  </button>
                  <span>Attempt {deliveryIndex}</span>
                  <span>{String(delivery.delivered_at)}</span>
                  <span>{delivery.status}</span>
                </li>
              );
            })}
          </ol>
        </div>
      </div>
    </div>
  );
};

const EventAttemptDetails = ({
  navigateEvent,
}: {
  navigateEvent: (path: string, params?: any) => void;
}) => {
  const { event_id: eventId } = useParams();
  const { deliveryIndex, delivery } = useLocation().state as {
    deliveryIndex: number;
    delivery: Delivery;
  };

  return (
    <div className="drawer">
      <div className="drawer__header">
        <div className="drawer__header-tabs">
          <Button minimal onClick={() => navigateEvent(`/${eventId}/attempts`)}>
            <PreviousIcon />
            Attempts {deliveryIndex}
          </Button>
        </div>

        <div>
          <Button minimal onClick={() => navigateEvent("/")}>
            <CloseIcon />
          </Button>
        </div>
      </div>

      <div className="drawer__body">
        <div className="event-attempt-details">
          <pre>{JSON.stringify(delivery, null, 2)}</pre>
        </div>
      </div>
    </div>
  );
};
