import useSWR from "swr";
import { useState } from "react";
import Table from "../../common/Table/Table";
import "./DestinationList.scss";
import SearchInput from "../../common/SearchInput/SearchInput";
import Button from "../../common/Button/Button";
import { AddIcon, FilterIcon } from "../../common/Icons";
import destinationTypes from "../../destination-types";

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

  if (!destinations) {
    return <div>Loading...</div>;
  }

  const table_columns = [
    { header: "Type", minWidth: 100, relativeWidth: 15 },
    { header: "Target", minWidth: 150, relativeWidth: 25 },
    { header: "Topics", minWidth: 150, relativeWidth: 20 },
    { header: "Status", minWidth: 100, relativeWidth: 10 },
    { header: "Success Rate", minWidth: 100, relativeWidth: 10 },
    { header: "Events (24h)", minWidth: 100, relativeWidth: 15 },
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
          <div style={{ minWidth: "20px", width: "20px", display: "flex" }}>
            {destinationTypes[destination.type].icon}
          </div>
          {destinationTypes[destination.type].label}
        </>,
        destination.config[destinationTypes[destination.type].target],
        destination.topics.join(", "),
        destination.disabled_at ? "Disabled" : "Active",
        "N/A", // TODO: Replace with actual success rate data
        "N/A", // TODO: Replace with actual events count
      ],
      link: `/${destination.id}`,
    })) || [];

  return (
    <div className="destination-list">
      <div className="destination-list__header">
        <h1 className="title-3xl">Event Destinations</h1>
        <div className="destination-list__actions">
          <SearchInput
            value={searchTerm}
            onChange={setSearchTerm}
            placeholder="Search by type, target, or topic..."
          />
          <Button onClick={console.log}>
            <FilterIcon /> Status
          </Button>
          <Button primary to="/new">
            <AddIcon /> Add Destination
          </Button>
        </div>
      </div>
      <Table columns={table_columns} rows={table_rows} />
    </div>
  );
};

export default DestinationList;
