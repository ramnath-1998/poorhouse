import React from 'react'
import Container from 'react-bootstrap/Container';
import Nav from 'react-bootstrap/Nav';
import Navbar from 'react-bootstrap/Navbar';

export const BlogNavbar = () => {
  return (
    <>
    <Navbar bg="dark" data-bs-theme="dark">
      <Container>
        <Navbar.Brand>PoorHouse</Navbar.Brand>
        <Nav className="me-auto">
          <Nav.Link href="/home">Home</Nav.Link>
          <Nav.Link href="/posts">Posts</Nav.Link>
          <Nav.Link href="/about">About</Nav.Link>
        </Nav>
      </Container>
    </Navbar>
  </>  )
}
