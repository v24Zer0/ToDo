import React from "react";
import { Text, View } from "react-native";
import List from "../models/list";

const ListComponent: React.FC<List> = ({id, name, user_id}) => {
    return (
        <View>
            <Text>
                {name}
            </Text>
        </View>
    );
}

export default ListComponent