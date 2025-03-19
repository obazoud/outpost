import { Route, Routes, useNavigate, useParams } from "react-router-dom";
import Button from "../../../common/Button/Button";
import { CloseIcon } from "../../../common/Icons";

const EventDetails = () => {
  const { destination_id: destinationId, event_id: eventId } = useParams();
  const navigate = useNavigate();

  const event = {
    metadata: {
      hello: "world",
    },
    data: {
      hello: "world",
    },
  };

  return (
    <div className="drawer">
      <div className="event-details">
        <div className="event-details__header">
          <div className="event-details__header-tabs">
            <Button to={`/destinations/${destinationId}/events/${eventId}`}>
              Event
            </Button>
            <Button
              to={`/destinations/${destinationId}/events/${eventId}/attempts`}
            >
              Attempts
            </Button>
          </div>

          <div>
            <Button
              minimal
              onClick={() => navigate(`/destinations/${destinationId}/events`)}
            >
              <CloseIcon />
            </Button>
          </div>
        </div>

        <div className="event-details__body">
          <Routes>
            <Route path="/" element={<EventData event={event} />} />
            <Route path="/attempts" element={<EventAttempts />} />
          </Routes>
        </div>
      </div>
    </div>
  );
};

export default EventDetails;

const EventData = ({ event }: { event: any }) => {
  return (
    <div className="event-data">
      <pre>{JSON.stringify(event, null, 2)}</pre>
    </div>
  );
};

const EventAttempts = () => {
  return (
    <div className="event-attempts">
      <div>Attempts</div>
    </div>
  );
};
