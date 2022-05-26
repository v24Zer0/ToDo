import React from "react";
import RootStackParamList from "./rootStackParamList";
import { createStackNavigator } from "@react-navigation/stack";
import LoginScreen from "./login_screen";
import ListScreen from "./list_screen";
import ItemScreen from "./item_screen";
import UserScreen from "./user_screen";
import SignupScreen from "./signup_screen";
import UserButton from "../components/user_button";

const Stack = createStackNavigator<RootStackParamList>();

const RootStack = () => {
    return (
        <Stack.Navigator>
            <Stack.Screen name="Login" component={LoginScreen} />
            <Stack.Screen name="Signup" component={SignupScreen} />
            <Stack.Screen name="List" component={ListScreen} 
                options={({ navigation }) => ({ 
                    headerRight: () => <UserButton navigate={() => navigation.navigate("User", { user: { id: "", username: "user1" }})} /> 
                })} 
            />
            <Stack.Screen name="Item" component={ItemScreen} 
                options={({ navigation, route }) => ({ 
                    title: route.params.list.name, 
                    headerRight: () => <UserButton navigate={() => navigation.navigate("User", { user: { id: "", username: "user1" }})} /> 
                })} 
            />
            <Stack.Screen name="User" component={UserScreen} 
                options={({ title: "Update details" })} 
            />
        </Stack.Navigator>
    );
}

export default RootStack;