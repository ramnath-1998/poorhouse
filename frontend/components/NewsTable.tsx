import React, { useEffect, useState } from 'react'
import Table from 'react-bootstrap/Table';

export const NewsTable = (props) => {

  const [webSocketUrl,setWebSocketUrl] = useState(props.url);

  useEffect(() => {

    const websocket = new WebSocket(webSocketUrl);

    websocket.onopen = () => {

      console.log("Connection established");
      websocket.send("Hi Server!")

    };

    websocket.addEventListener("message", (event) => {
      console.log("Message from server ", event.data);
    });

    websocket.onclose = () => {
      console.log("Connection closed!");

    };



    return () => {
      websocket.close();
    }
  }, [])
  return (
    <div>
    <Table striped>
    <tbody>
      <tr>
      <td>Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Aenean commodo ligula eget dolor.</td>
      </tr>
      <tr>
      <td>Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Aenean commodo ligula eget dolor.</td>
      </tr>
      <tr>
      <td>Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Aenean commodo ligula eget dolor.</td>
      </tr>
    </tbody>
  </Table>
  </div>
  )
}
