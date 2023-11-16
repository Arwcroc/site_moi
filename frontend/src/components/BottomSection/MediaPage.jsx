import React from 'react';
import { useState, useEffect } from 'react';
import Box from '@mui/material/Box';
import Grid from '@mui/material/Grid';
import Typography from '@mui/material/Typography';
import Avatar from '@mui/material/Avatar';
import Link from '@mui/material/Link';
import CircularProgress from '@mui/material/CircularProgress';
import GroupTwoToneIcon from '@mui/icons-material/GroupTwoTone';


// extension vscode : prettier - code formatter

const TwitchElement = (props) => {
	const [twitchData, setTwitchData] = useState();

	const fetchTwitch = () => {
		if (twitchData !== undefined) return;
		console.log("fetching api twitch " + props.username)
		fetch('http://localhost:8090/twitchapi?user_id=' + props.username).then( async (response) => {
			if (response.status < 200 || response.status >= 400) {
				console.error("Error on fetch api twitch " + props.username)
			};
			let body = await response.json();
			console.log("got api twitch " + props.username)
			setTwitchData(body);
		})
	}
	useEffect(() => {
		fetchTwitch()
		// eslint-disable-next-line
	}, [setTwitchData]);

	return (
		<Grid item xs={8} md={6}>
		{
			twitchData ? (
			<>
			<Link href={"https://www.twitch.tv/" + props.username} color="inherit" underline='none' target="_blank">
				<Box className="App__WebContainer__BottomSection__MediaPage__TwitchPlace__Box">
					<Box className="App__WebContainer__BottomSection__MediaPage__TwitchPlace__Box__Pics">	
						<Avatar src={twitchData.user.profile_image_url}>Img</Avatar> 
					</Box>
					<Box className="App__WebContainer__BottomSection__MediaPage__TwitchPlace__Box__Info">	
						<Typography gutterBottom variant="subtitle1" component="div">
							{twitchData.user.display_name}
						</Typography>
						<Box className="App__WebContainer__BottomSection__MediaPage__TwitchPlace__Box__Info_Live">
							{twitchData.has_stream === false ? <>{"VISIT STREAMER PAGE"} </>: <>{twitchData.stream.game_name + " : " + twitchData.stream.viewer_count + " "} <GroupTwoToneIcon/></>}
						</Box>
					</Box>		
				</Box>
			</Link>
			</>
			): (
				<>
					<Box>
						<CircularProgress />
					</Box>
				</>
			)	
		}			
		</Grid>
	)
}

const MediaPage = () => {

	return (
		<div className="App__WebContainer__BottomSection__MediaPage">
			<div className="App__WebContainer__BottomSection__MediaPage__TwitchPlace">
				<Grid 
				container
				columnSpacing={8}
				rowSpacing={8}
				direction="row"
				justifyContent="center"
				alignItems="center"
				>
					<TwitchElement username="joueur_du_grenier" />
					<TwitchElement username="zerator" />
					<TwitchElement username="antoinedaniel" />
					<TwitchElement username="ponce" />
				</Grid>
			</div>
			<div className="App__WebContainer__BottomSection__MediaPage__YoutubePlace">
				<Grid 
				container
				columnSpacing={8}
				rowSpacing={6}
				direction="row"
				justifyContent="center"
				alignItems="center"
				>
					<TwitchElement username="joueur_du_grenier" />
					<TwitchElement username="zerator" />
					<TwitchElement username="antoinedaniel" />
					<TwitchElement username="ponce" />
				</Grid>
			</div>
		</div>
	);
}

export default MediaPage;