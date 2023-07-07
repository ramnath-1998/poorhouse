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
    <Row>
      <Col md={6}>
      <NewsTable></NewsTable>
      </Col>

      <Col md={6}>
      <NewsTable></NewsTable>
      </Col>
    </Row>
    <NewsPopUp url="https://www.news18.com/elections/west-bengal-panchayat-elections-july-8-central-forces-tmc-bjp-isf-rural-polls-8280301.html"></NewsPopUp>
    </>

  )
}
