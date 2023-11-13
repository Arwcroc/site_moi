import React from 'react';
import { Box, TextField } from '@mui/material';

const ContactPage = () => {
	return (
		<Box className="App__WebContainer__BottomSection__ContactPage__Main">
			Contactez moi :
			<TextField id="outlined-basic" label="Full Name" variant="outlined" margin="dense" className="App__WebContainer__BottomSection__ContactPage__Main__Name">
			</TextField>
			<TextField id="outlined-basic" label="Company" variant="outlined" margin="dense" className="App__WebContainer__BottomSection__ContactPage__Main__Name">
			</TextField>
			<TextField id="outlined-basic" label="E-Mail" variant="outlined" margin="dense" className="App__WebContainer__BottomSection__ContactPage__Main__Name">
			</TextField>
			<TextField id="outlined-basic" label="Message" multiline maxRows={4} margin="dense" className="App__WebContainer__BottomSection__ContactPage__Main__Name">
			</TextField>
		</Box>
	);
}

export default ContactPage;