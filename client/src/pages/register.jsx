import React, {useState} from "react";
import { Form, Button } from "react-bootstrap";
import {useDispatch} from 'react-redux'
import {registerAPI }from "../redux/action/authAction"

function Register() {

 const [fullName, setFullname] = useState("")
 const [address, setAdress] = useState("")
 const [email, setEmail] = useState("")
 const [password, setPassword] = useState("")


const dispatch = useDispatch()

 const handleSubmit = (e) => {
    e.preventDefault()

    const data = {
        fullname: fullName,
        address: address,
        email: email,
        password: password,

    };
console.log(data);

dispatch(registerAPI)

}

  return (
    <div>
      <Form onSubmit={handleSubmit}>
        <Form.Group className="mb-3" controlId="formBasicFullname">
          <Form.Label>Fullname</Form.Label>
          <Form.Control onChange={(e) => setFullname(e.target.value)} type="fullname" placeholder="fullname" />
        </Form.Group>
        
        <Form.Group className="mb-3" controlId="formBasicAddress">
          <Form.Label>Address</Form.Label>
          <Form.Control onChange={(e) => setAdress(e.target.value)}  type="address" placeholder="address" />
        </Form.Group>
        
        <Form.Group className="mb-3" controlId="formBasicEmail">
          <Form.Label>Email address</Form.Label>
          <Form.Control onChange={(e) => setEmail(e.target.value)}  type="email" placeholder="Enter email" />
        </Form.Group>

        <Form.Group className="mb-3" controlId="formBasicPassword">
          <Form.Label>Password</Form.Label>
          <Form.Control onChange={(e) => setPassword(e.target.value)}  type="password" placeholder="Password" />
        </Form.Group>
        <Button variant="primary" type="submit">
          Submit
        </Button>
      </Form>
    </div>
  );
}

export default Register;
