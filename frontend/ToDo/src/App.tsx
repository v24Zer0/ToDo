/**
 * Sample React Native App
 * https://github.com/facebook/react-native
 *
 * Generated with the TypeScript template
 * https://github.com/react-native-community/react-native-template-typescript
 *
 * @format
 */

import React from 'react';
import { NavigationContainer } from '@react-navigation/native';

import MyStack from './stack';

const App = () => {
	// const backgroundStyle = {
	// 	backgroundColor: isDarkMode ? Colors.darker : Colors.lighter,
	// };

	return (
		<NavigationContainer>
			<MyStack />
		</NavigationContainer>
	);
};

export default App;
