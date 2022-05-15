import React from "react";
import LoginScreen from "./loginScreen";
import RootStackParamList from "./rootStackParamList";
import { createStackNavigator } from "@react-navigation/stack";
import ListScreen from "./listScreen";
import ItemScreen from "./itemScreen";

const Stack = createStackNavigator<RootStackParamList>();

const RootStack = () => {
    return (
        <Stack.Navigator>
            <Stack.Screen name="Login" component={LoginScreen} />
            {/* <Stack.Screen name="Signup" component={Signup} /> */}
            <Stack.Screen name="List" component={ListScreen} />
            <Stack.Screen name="Item" component={ItemScreen} />
        </Stack.Navigator>
    );
}

export default RootStack;