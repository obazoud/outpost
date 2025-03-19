import { useCallback, useMemo, useState } from "react";
import Badge from "../../../common/Badge/Badge";
// import Button from "../../../common/Button/Button";
// import SearchInput from "../../../common/SearchInput/SearchInput";
import "./Events.scss";
import Table from "../../../common/Table/Table";
import { EventListResponse } from "../../../typings/Event";
import useSWR from "swr";
import Dropdown from "../../../common/Dropdown/Dropdown";
import { CalendarIcon, FilterIcon, RefreshIcon } from "../../../common/Icons";
import { Checkbox } from "../../../common/Checkbox/Checkbox";
import {
  Route,
  Routes,
  useNavigate,
  useSearchParams,
  Outlet,
  useParams,
} from "react-router-dom";
import Button from "../../../common/Button/Button";
import CONFIGS from "../../../config";
import EventDetails from "./EventDetails";

const Events = ({
  destination,
  navigateEvent,
}: {
  destination: any;
  navigateEvent: (path: string, state?: any) => void;
}) => {
  const { status, topics, urlSearchParams } = useEventFilter();
  const [timeRange, setTimeRange] = useState("24h");
  const { event_id: eventId } = useParams();

  const queryUrl = useMemo(() => {
    const now = new Date();
    let start;
    switch (timeRange) {
      case "7d":
        start = new Date(now.getTime() - 7 * 24 * 60 * 60 * 1000);
        break;
      case "30d":
        start = new Date(now.getTime() - 30 * 24 * 60 * 60 * 1000);
        break;
      default: // 24h
        start = new Date(now.getTime() - 24 * 60 * 60 * 1000);
    }
    return `destinations/${
      destination.id
    }/events?start=${start.toISOString()}&${urlSearchParams}`;
  }, [destination.id, timeRange, urlSearchParams]);

  const {
    data: eventsList,
    mutate,
    isValidating,
  } = useSWR<EventListResponse>(queryUrl);

  const topicsList = CONFIGS.TOPICS.split(",");

  const table_rows = eventsList
    ? eventsList.data.map((event) => ({
        id: event.id,
        entries: [
          <span className="mono-s">
            {new Date(event.time).toLocaleString("en-US", {
              year: "numeric",
              month: "short",
              day: "numeric",
              hour: "numeric",
              minute: "2-digit",
              hour12: true,
            })}
          </span>,
          <span className="mono-s">
            {event.status === "success" ? (
              <Badge text="Successful" success />
            ) : event.status === "failed" ? (
              <Badge text="Failed" danger />
            ) : (
              <Badge text="Pending" />
            )}
          </span>,
          <span className="mono-s">{event.topic}</span>,
          <span className="mono-s">{event.id}</span>,
        ],
        onClick: () => navigateEvent(`/${event.id}`),
      }))
    : [];

  return (
    <div className="destination-events">
      <div className="destination-events__header">
        <h2 className="title-l">
          Events <Badge text={eventsList ? eventsList.data.length + "" : "0"} />
        </h2>
        <div className="destination-events__header-filters">
          {/* <SearchInput
            value={search}
            onChange={setSearch}
            placeholder="Filter by ID"
          /> */}
          <Dropdown
            trigger_icon={<CalendarIcon />}
            trigger={`Last ${timeRange}`}
          >
            <div className="dropdown-item">
              <Checkbox
                label="Last 24h"
                checked={timeRange === "24h"}
                onChange={() => setTimeRange("24h")}
              />
            </div>
            <div className="dropdown-item">
              <Checkbox
                label="Last 7d"
                checked={timeRange === "7d"}
                onChange={() => setTimeRange("7d")}
              />
            </div>
            <div className="dropdown-item">
              <Checkbox
                label="Last 30d"
                checked={timeRange === "30d"}
                onChange={() => setTimeRange("30d")}
              />
            </div>
          </Dropdown>

          <Dropdown trigger_icon={<FilterIcon />} trigger="Status">
            <div className="dropdown-item">
              <Checkbox
                label="Success"
                checked={status.value === "success"}
                onChange={() =>
                  status.value === "success"
                    ? status.set("")
                    : status.set("success")
                }
              />
            </div>
            <div className="dropdown-item">
              <Checkbox
                label="Failed"
                checked={status.value === "failed"}
                onChange={() =>
                  status.value === "failed"
                    ? status.set("")
                    : status.set("failed")
                }
              />
            </div>
          </Dropdown>

          <Dropdown trigger_icon={<FilterIcon />} trigger="Topics">
            {topicsList.map((topic) => (
              <div key={topic} className="dropdown-item">
                <Checkbox
                  label={topic}
                  checked={topics.value.includes(topic)}
                  onChange={() => topics.toggle(topic)}
                />
              </div>
            ))}
          </Dropdown>

          <Button
            onClick={() => mutate()}
            disabled={isValidating}
            loading={isValidating}
          >
            <RefreshIcon />
            Refresh
          </Button>
        </div>
      </div>

      <div
        className={`destination-events__table ${eventId ? "with-drawer" : ""}`}
      >
        <Table
          columns={[
            {
              header: "Timestamp",
            },
            {
              header: "Status",
            },
            {
              header: "Topic",
            },
            {
              header: "Message ID",
            },
          ]}
          rows={table_rows}
          footer_label="events"
        />

        <Outlet />
      </div>
    </div>
  );
};

export default Events;

const useEventFilter = () => {
  const [searchParams, setSearchParams] = useSearchParams();

  const handleFilterChange = (key: string, value: string | null) => {
    setSearchParams((prev) => {
      const params = new URLSearchParams(prev);
      if (value) {
        params.set(key, value);
      } else {
        params.delete(key);
      }
      return params;
    });
  };

  const status = {
    value: searchParams.get("status") || "",
    set: (value: string) => handleFilterChange("status", value || null),
  };

  const topics = {
    value: searchParams.getAll("topic"),
    set: (value: string[]) => {
      setSearchParams((prev) => {
        const params = new URLSearchParams(prev);
        params.delete("topic");
        value.forEach((v) => params.append("topic", v));
        return params;
      });
    },
    toggle: (topic: string) => {
      const currentTopics = searchParams.getAll("topic");
      const newTopics = currentTopics.includes(topic)
        ? currentTopics.filter((t) => t !== topic)
        : [...currentTopics, topic];
      setSearchParams((prev) => {
        const params = new URLSearchParams(prev);
        params.delete("topic");
        newTopics.forEach((t) => params.append("topic", t));
        return params;
      });
    },
  };

  const urlSearchParams = useMemo(() => {
    return searchParams.toString();
  }, [searchParams]);

  return {
    status,
    topics,
    urlSearchParams,
  };
};

export const EventRoutes = ({ destination }: { destination: any }) => {
  const { urlSearchParams } = useEventFilter();
  const navigate = useNavigate();

  const navigateEvent = useCallback(
    (path: string, state?: any) => {
      navigate(
        `/destinations/${destination.id}/events${path}?${urlSearchParams}`,
        { state }
      );
    },
    [navigate, urlSearchParams]
  );

  return (
    <Routes>
      <Route
        path="/"
        element={
          <Events destination={destination} navigateEvent={navigateEvent} />
        }
      >
        <Route
          path=":event_id/*"
          element={<EventDetails navigateEvent={navigateEvent} />}
        />
      </Route>
    </Routes>
  );
};
