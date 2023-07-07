import React from 'react'
import ReactDOM from 'react-dom/client';

export const Application = () => {
  return (
    <div>Application</div>
  )
}
const domContainer = document.querySelector('#application');
const root = ReactDOM.createRoot(domContainer!);
root.render(<Application></Application>);
