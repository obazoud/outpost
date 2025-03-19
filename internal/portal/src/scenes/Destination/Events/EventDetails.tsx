import { Route, Routes, useLocation, useParams } from "react-router-dom";
import Button from "../../../common/Button/Button";
import { CloseIcon, NextIcon, PreviousIcon } from "../../../common/Icons";
import useSWR from "swr";
import { Event, Delivery } from "../../../typings/Event";
import Badge from "../../../common/Badge/Badge";

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

  if (!event) {
    return <div>Error loading event</div>;
  }

  return (
    <div className="drawer">
      <div className="drawer__header">
        <div className="drawer__header-tabs">
          <nav className="tabs">
            <button
              type="button"
              onClick={() => navigateEvent(`/${eventId}`)}
              className={`tab tab--active`}
            >
              Event
            </button>
            <button
              type="button"
              onClick={() => navigateEvent(`/${eventId}/attempts`)}
              className={`tab`}
            >
              Attempts
            </button>
          </nav>
        </div>

        <div className="drawer__header-actions">
          <Badge
            text={event.status === "success" ? "Successful" : "Failed"}
            success={event.status === "success"}
            danger={event.status === "failed"}
          />

          <Button minimal onClick={() => navigateEvent("/")}>
            <CloseIcon />
          </Button>
        </div>
      </div>

      <div className="drawer__body">
        <div className="event-data">
          <div className="event-data__overview">
            <h3 className="title-m">Overview</h3>
            <div>
              <div>ID: {event.id}</div>
              <div>
                Topic: <code>{event.topic}</code>
              </div>
              <div>Created at: {new Date(event.time).toLocaleString()}</div>
            </div>
          </div>

          <div className="event-data__data">
            <h3 className="title-m">Data</h3>
            <pre>{JSON.stringify(event.data, null, 2)}</pre>
          </div>

          <div className="event-data__metadata">
            <h3 className="title-m">Metadata</h3>
            <pre>{JSON.stringify(event.metadata, null, 2)}</pre>
          </div>
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
          <nav className="tabs">
            <button
              type="button"
              onClick={() => navigateEvent(`/${eventId}`)}
              className={`tab`}
            >
              Event
            </button>
            <button
              type="button"
              onClick={() => navigateEvent(`/${eventId}/attempts`)}
              className={`tab tab--active`}
            >
              Attempts
            </button>
          </nav>
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

                  <span className="event-attempts__list-item-content">
                    <span>
                      <span className="title-m">Attempt {deliveryIndex}</span>
                      <span className="mono-s">
                        {new Date(delivery.delivered_at).toLocaleString(
                          "en-US",
                          {
                            month: "short",
                            day: "numeric",
                            hour: "numeric",
                            minute: "2-digit",
                            hour12: true,
                          }
                        )}
                      </span>
                    </span>
                    <span>
                      <Badge
                        text={
                          delivery.status === "success"
                            ? "Successful"
                            : "Failed"
                        }
                        success={delivery.status === "success"}
                        danger={delivery.status === "failed"}
                      />
                      {/* TODO: use RightArrowIcon */}
                      <NextIcon />
                    </span>
                  </span>
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

        <div className="drawer__header-actions">
          <Badge
            text={delivery.status === "success" ? "Successful" : "Failed"}
            success={delivery.status === "success"}
            danger={delivery.status === "failed"}
          />

          <Button minimal onClick={() => navigateEvent("/")}>
            <CloseIcon />
          </Button>
        </div>
      </div>

      <div className="drawer__body">
        <div className="event-attempt-details">
          <div className="event-attempt-details__overview">
            <h3 className="title-m">Overview</h3>

            <div>ID: {delivery.id}</div>
            <div>
              Delivered at {new Date(delivery.delivered_at).toLocaleString()}
            </div>
          </div>
          {/* Currently, we're not storing the exact message data sent. */}
          {/* <div className="event-attempt-details__message">
            <h3 className="title-s">Message</h3>
          </div> */}
          <div className="event-attempt-details__result">
            <h3 className="title-m">Response</h3>
            <pre>{JSON.stringify(delivery.response_data, null, 2)}</pre>
          </div>
        </div>
      </div>
    </div>
  );
};
