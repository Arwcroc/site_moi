import React from 'react';
import { Box, TextField, Typography } from '@mui/material';

const ContactPage = () => {
	return (
		<div className="App__WebContainer__BottomSection__ContactPage__Main">
			<Box className="App__WebContainer__BottomSection__ContactPage__Info">
				<Typography>Contactez moi :</Typography>
				<TextField id="outlined-basic" fullWidth label="Full Name" variant="outlined" margin="dense" className="App__WebContainer__BottomSection__ContactPage__Main__Name">
				</TextField>
				<TextField id="outlined-basic" fullWidth label="Company" variant="outlined" margin="dense" className="App__WebContainer__BottomSection__ContactPage__Main__Name">
				</TextField>
				<TextField id="outlined-basic" fullWidth label="E-Mail" variant="outlined" margin="dense" className="App__WebContainer__BottomSection__ContactPage__Main__Name">
				</TextField>
				<TextField id="outlined-basic" fullWidth label="Message" multiline maxRows={4} margin="dense" className="App__WebContainer__BottomSection__ContactPage__Main__Name">
				</TextField>
			</Box>
			<Box className="App__WebContainer__BottomSection__ContactPage__Send">
				<Typography>Send</Typography>
			</Box>
		</div>
	);
}

export default ContactPage;