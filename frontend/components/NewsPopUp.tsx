import React, { useState } from 'react';
import { Modal, Button } from 'react-bootstrap';
export const NewsPopUp = (props) => {

    const [showPopup, setShowPopup] = useState(false);

    const openPopup = () => {
        setShowPopup(true);
    };

    const closePopup = () => {
        setShowPopup(false);
    };


    return (
        <div>
        <Button onClick={openPopup}>Open Popup</Button>
        <Modal show={showPopup} onHide={closePopup}>
        <Modal.Header closeButton>
        <Modal.Title>Popup Window</Modal.Title>
        </Modal.Header>
        <Modal.Body>
        <iframe src={props.url} width="100%" height="400px" title="Website Modal" />
        </Modal.Body>
        <Modal.Footer>
        <Button variant="secondary" onClick={closePopup}>
        Close
        </Button>
        </Modal.Footer>
        </Modal>
        </div>
        )
    }

