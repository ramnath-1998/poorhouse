import React from 'react'
import ReactDOM from 'react-dom/client';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import {Home} from './pages/Home';
import {About} from './pages/About';
import {Posts} from './pages/Posts';

const domContainer = document.querySelector('#application');
const root = ReactDOM.createRoot(domContainer!);
root.render(

<BrowserRouter>
<Routes>
  <Route index element={<Home></Home>}></Route>
  <Route path="/home" element={<Home></Home>}></Route>
  <Route path="/about" element = {<About></About>}></Route>
  <Route path="/posts" element = {<Posts></Posts>}></Route>
</Routes>
</BrowserRouter>);
