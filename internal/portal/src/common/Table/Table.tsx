import { Link } from "react-router-dom";
import "./Table.scss";

interface Column {
  header: string;
  maxWidth?: number;
  minWidth?: number;
  relativeWidth: number;
}

interface TableRow {
  id: string;
  entries: (string | React.ReactNode)[];
  link?: string;
  onClick?: () => void;
}

interface TableProps {
  columns: Column[];
  rows: TableRow[];
}

const Table: React.FC<TableProps> = ({ columns, rows }) => {
  const handle_row_click = (row: TableRow) => {
    if (row.onClick) {
      row.onClick();
    }
  };

  return (
    <div className="table-container">
      <table className="table">
        <thead>
          <tr>
            {columns.map((column, index) => (
              <th
                key={`header-${index}`}
                className="subtitle-xs"
                style={{
                  maxWidth: column.maxWidth,
                  minWidth: column.minWidth,
                  width: `${column.relativeWidth}%`,
                }}
              >
                {column.header}
              </th>
            ))}
          </tr>
        </thead>
        <tbody>
          {rows.map((row) => (
            <tr key={row.id} onClick={() => handle_row_click(row)}>
              {row.entries.map((entry, index) => (
                <td
                  key={`${row.id}-${index}`}
                  style={{
                    maxWidth: columns[index]?.maxWidth,
                    minWidth: columns[index]?.minWidth,
                    width: `${columns[index]?.relativeWidth}%`,
                  }}
                >
                  {row.link ? <Link to={row.link}>{entry}</Link> : entry}
                </td>
              ))}
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default Table;
