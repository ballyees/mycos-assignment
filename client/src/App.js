import './App.css';
import { Routes, Route } from "react-router-dom";
import { Home } from './components';
import { Container, Row, Col } from 'react-bootstrap'


function App() {
  return (
    <div className="App">
      <Container fluid="md" style={{paddingTop: 20}}>
        <Row><Col></Col></Row>
        <Row>
          <Col>
            <Routes>
              <Route path="/" element={<Home />} />
            </Routes>
          </Col>
        </Row>
      </Container>
    </div>
  );
}

export default App;
