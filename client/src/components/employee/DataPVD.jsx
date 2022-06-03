import { useEffect, useState } from 'react';
import Table from 'react-bootstrap/Table';
import Config from '../../config';

const columns = [
  { feild: 'ID', displayName: 'ID' },
  { feild: 'FirstName', displayName: 'First Name' },
  { feild: 'LastName', displayName: 'Last Name' },
  { feild: 'BirthDate', displayName: 'Birth Date' },
  { feild: 'StartDate', displayName: 'Start Date' },
  { feild: 'Salary', displayName: 'Salary', useComma: true },
  { feild: 'PvdFundRate', displayName: 'PVD Fund Rate' },
  { feild: 'TotalPVD', displayName: 'Total PVD', useComma: true },
]

function DataEmployeePVD() {
  const [data, setData] = useState([]);
  const [isLoading, setIsLoading] = useState(true);
  useEffect(() => {
    fetch(Config.getConfig().pvdURL)
      .then(res => res.json())
      .then(res => {
        res = res.map(data => {
          data = ({ ...(data?.Employee || {}), ...data })
          delete data?.Employee
          return data
        })
        setData(res)
        setIsLoading(false)
      });
  }, []) // run only once when the component is mounted
  return (
    <>
      <h1>{isLoading ? "Loading ..." : null}</h1>
      <Table responsive striped bordered>
        <thead>
          <tr>
            {columns.map(column => <th key={column.feild}>{column.displayName}</th>)}
          </tr>
        </thead>
        <tbody>
          {
            data.map(row => (
              <tr key={row.ID}>
                {columns.map(column => <td key={column.feild}>{column.useComma ? numberWithCommas(row[column.feild]) : row[column.feild]}</td>)}
              </tr>
            )
            )
          }
        </tbody>
      </Table>
    </>
  )
}

function numberWithCommas(number) {
  return number.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ",");
}

export default DataEmployeePVD;