import React from 'react'
import { Link } from 'react-router-dom'
import { BlogNavbar } from '../components/BlogNavbar'
import { NewsTable } from '../components/NewsTable'
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import { NewsPopUp } from '../components/NewsPopUp';

export const Home = () => {
  return (
    <><BlogNavbar></BlogNavbar>
    <NewsTable url="ws://localhost:3000/ws"></NewsTable>
    </>

  )
}
