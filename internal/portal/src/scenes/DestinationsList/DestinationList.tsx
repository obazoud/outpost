import useSWR from "swr";

interface Destination {
  id: string;
  name: string;
  url: string;
}

const DestinationList: React.FC = () => {
  const { data: destinations } = useSWR<Destination[]>("destinations");

  if (!destinations) {
    return <div>Loading...</div>;
  }

  return (
    <div className="destination-list">
      <ul>
        {destinations.map((destination) => (
          <li key={destination.id}>
            <a href={destination.url} target="_blank" rel="noopener noreferrer">
              {destination.name}
            </a>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default DestinationList;
