import { useState } from "react";
import Badge from "../../../common/Badge/Badge";
import Button from "../../../common/Button/Button";
import SearchInput from "../../../common/SearchInput/SearchInput";
import "./Events.scss";
import Table from "../../../common/Table/Table";

const Events = ({ destination }: { destination: any }) => {
  const [search, setSearch] = useState("");

  return (
    <div className="destination-events">
      <div className="destination-events__header">
        <h2 className="title-l">
          Events [TODO] <Badge text="100" />
        </h2>
        <div className="destination-events__header-filters">
          <SearchInput
            value={search}
            onChange={setSearch}
            placeholder="Filter by ID"
          />
          <Button>Last 24 Hours</Button>
          <Button>Status</Button>
          <Button>Topics</Button>
          <Button>Refresh</Button>
        </div>
      </div>

      <div className="destination-events__table">
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
          rows={[]}
          onClick={console.log}
        />
      </div>
    </div>
  );
};

export default Events;
