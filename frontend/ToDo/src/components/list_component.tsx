import React from "react";
import { Button, Pressable, Text, View } from "react-native";
import List from "../models/list";

interface ListProps {
    list: List;
    setModalVisible: React.Dispatch<React.SetStateAction<boolean>>;
    setModalList: React.Dispatch<React.SetStateAction<List>>;
}

// Add Pressable component to trigger modal display
// onPress triggers ItemList
// onLongPress triggers Modal display
const ListComponent: React.FC<ListProps> = ({ list, setModalVisible, setModalList }) => {
    return (
        <View>
            <Pressable onLongPress={() => {
                    setModalVisible(true);
                    setModalList(list);
                }}
            >
                <Text>
                    {list.name}
                </Text>
            </Pressable>
        </View>
    );
}

export default ListComponent