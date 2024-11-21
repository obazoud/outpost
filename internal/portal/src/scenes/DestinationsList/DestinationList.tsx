import useSWR from "swr";
import { useState } from "react";
import Table from "../../common/Table/Table";
import "./DestinationList.scss";
import SearchInput from "../../common/SearchInput/SearchInput";
import Button from "../../common/Button/Button";
import { AddIcon, FilterIcon } from "../../common/Icons";
import destinationTypes from "../../destination-types";
import Badge from "../../common/Badge/Badge";
import Tooltip from "../../common/Tooltip/Tooltip";

interface Destination {
  id: string;
  type: "webhooks";
  config: {
    url: string;
  };
  topics: string[];
  disabled_at: string | null;
  // Add other fields that might be needed for the new columns
}

const DestinationList: React.FC = () => {
  const { data: destinations } = useSWR<Destination[]>("destinations");
  const [searchTerm, setSearchTerm] = useState("");

  // TODO: Add loading state
  if (!destinations) {
    return <div>Loading...</div>;
  }

  const table_columns = [
    { header: "Type", width: 160 },
    { header: "Target" },
    { header: "Topics", width: 120 },
    { header: "Status", width: 120 },
    { header: "Success Rate", width: 120 },
    { header: "Events (24h)", width: 120 },
  ];

  const filtered_destinations = destinations?.filter((destination) => {
    const search_value = searchTerm.toLowerCase();
    return (
      destination.type.toLowerCase().includes(search_value) ||
      destination.config[destinationTypes[destination.type].target]
        .toLowerCase()
        .includes(search_value) ||
      destination.topics.some((topic) =>
        topic.toLowerCase().includes(search_value)
      )
    );
  });

  const table_rows =
    filtered_destinations?.map((destination) => ({
      id: destination.id,
      entries: [
        <>
          <div style={{ minWidth: "16px", width: "16px", display: "flex" }}>
            {destinationTypes[destination.type].icon}
          </div>
          <span className="subtitle-m">{destinationTypes[destination.type].label}</span>
        </>,
        <span className="muted-variant">
          {destination.config[destinationTypes[destination.type].target]}
        </span>,
        <Tooltip
          content={
            <div className="destination-list__topics-tooltip">
              {(destination.topics.length > 0 && destination.topics[0] === "*"
                ? TOPICS.split(",")
                : destination.topics
              )
                .slice(0, 9)
                .map((topic) => (
                  <Badge key={topic} text={topic.trim()} />
                ))}
              {(destination.topics[0] === "*"
                ? TOPICS.split(",").length
                : destination.topics.length) > 9 && (
                <span className="subtitle-s muted">
                  +{" "}
                  {(destination.topics[0] === "*"
                    ? TOPICS.split(",").length
                    : destination.topics.length) - 9}{" "}
                  more
                </span>
              )}
            </div>
          }
        >
          <span className="muted-variant">
            {destination.topics.length > 0 && destination.topics[0] === "*"
              ? "All"
              : destination.topics.length}
          </span>
        </Tooltip>,
        destination.disabled_at ? (
          <Badge text="Disabled" />
        ) : (
          <Badge text="Active" success />
        ),
        <span className="muted-variant">99.5%</span>, // TODO: Replace with actual success rate data
        <span className="muted-variant">100</span>, // TODO: Replace with actual events count
      ],
      link: `/destinations/${destination.id}`,
    })) || [];

  return (
    <div className="destination-list">
      <div className="destination-list__header">
        <span className="subtitle-s muted">&nbsp;</span>
        <h1 className="title-3xl">Event Destinations</h1>
        <div className="destination-list__actions">
          <SearchInput
            value={searchTerm}
            onChange={setSearchTerm}
            placeholder="Filter by type, target or topic"
          />
          {/* <Button onClick={console.log}>
            <FilterIcon /> Status
          </Button> */}
          <Button primary to="/new">
            <AddIcon /> Add Destination
          </Button>
        </div>
      </div>
      <Table
        columns={table_columns}
        rows={table_rows}
        footer_label="event destinations"
      />
    </div>
  );
};

export default DestinationList;
