import { useState, useContext } from "react";
import Button from "../../../common/Button/Button";
import TopicPicker from "../../../common/TopicPicker/TopicPicker";

import "./DestinationSettings.scss";
import { SaveIcon } from "../../../common/Icons";
import { ApiContext } from "../../../app";
import { mutate } from "swr";
import { showToast } from "../../../common/Toast/Toast";

const DestinationSettings = ({ destination }: { destination: any }) => {
  const [selectedTopics, setSelectedTopics] = useState(destination.topics);
  const apiClient = useContext(ApiContext);

  if (!apiClient) throw new Error("ApiContext not found");

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    apiClient
      .fetch(`destinations/${destination.id}`, {
        method: "PATCH",
        body: JSON.stringify({
          topics: selectedTopics,
        }),
        headers: {
          "Content-Type": "application/json",
        },
      })
      .then((data) => {
        showToast("success", "Destination event topics updated.");
        mutate(`destinations/${destination.id}`, data, false);
      })
      .catch((error) => {
        showToast(
          "error",
          `${error.message.charAt(0).toUpperCase() + error.message.slice(1)}`
        );
      });
  };

  return (
    <div className="destination-settings">
      <div className="destination-settings__topics">
        <h2 className="title-l">Event Topics</h2>
        <form onSubmit={handleSubmit}>
          <TopicPicker
            maxHeight="256px"
            selectedTopics={selectedTopics}
            onTopicsChange={(topics) => {
              setSelectedTopics(topics);
            }}
          />
          <Button
            className="destination-settings__submit-button"
            type="submit"
            disabled={
              JSON.stringify(selectedTopics) ===
                JSON.stringify(destination.topics) ||
              selectedTopics.length === 0
            }
            primary
          >
            <SaveIcon /> Save
          </Button>
        </form>
      </div>
      <hr />
      <div>
        <h2 className="title-l">Configuration & Credentials [TODO]</h2>
        <form>
          <p>TODO</p>
          <Button type="submit" disabled={true} primary>
            <SaveIcon /> Save
          </Button>
        </form>
      </div>
      <hr />
      <div>
        <h2 className="title-l">Disable event destination [TODO]</h2>
        <p className="body-m">
          Disabling an event destination will prevent it from receiving events.
        </p>
        <Button>Disable</Button>
      </div>
      <div>
        <h2 className="title-l">Delete event destination [TODO]</h2>
        <p className="body-m">
          Deleting an event destination is irreversible. All associated events
          will also be deleted.
        </p>
        <Button>Delete</Button>
      </div>
    </div>
  );
};

export default DestinationSettings;
