import React, { useEffect, useState } from "react";
import { Button, FlatList, ListRenderItem, Modal, View } from "react-native";
import List from "../models/list";
import ListComponent from "./list_component";
import ListCreateModal from "./list_create_modal";
import ListModal from "./list_modal";

const mockLists: List[] = [
    {
        id: "unique_list1",
        name: "list1",
        user_id: "user1"
    },
    {
        id: "unique_list2",
        name: "list2",
        user_id: "user2"
    }
];

const mockLists2: List[] = [
    {
        id: "unique_list13",
        name: "list3",
        user_id: "user1"
    },
    {
        id: "unique_list4",
        name: "list4",
        user_id: "user1"
    }
];

interface Props {
    navigate(list: List): void;
}

// Add Modal and pass state functions to ListComponent
const UserLists: React.FC<Props> = ({ navigate }) => {
    const [lists, setLists] = useState<List[]>(mockLists);
    const [update, setUpdate] = useState<boolean>(false);

    const [modalVisible, setModalVisible] = useState<boolean>(false);
    const [modalList, setModalList] = useState<List>({id: "", name: "", user_id: "user1"});

    const [createModalVisible, setCreateModalVisible] = useState<boolean>(false);

    useEffect(() => {
        update ? setLists(mockLists2) : setLists(mockLists)
    }, [update]);

    const renderItem: ListRenderItem<List> = ({ item }) => (
        <ListComponent list={item} setModalVisible={setModalVisible} setModalList={setModalList} navigate={navigate} />
    );

    return (
        <View>
            <Modal visible={modalVisible} onRequestClose={() => setModalVisible(false)}>
                <ListModal list={modalList} setUpdate={() => setUpdate(!update)} 
                    setModalVisible={() => setModalVisible(false)} 
                />
            </Modal>
            <Modal visible={createModalVisible} onRequestClose={() => setCreateModalVisible(false)}>
                <ListCreateModal setModalVisible={() => setCreateModalVisible(false)} />
            </Modal>
            <FlatList data={lists} renderItem={renderItem} keyExtractor={list => list.id}/>
            <Button title="Create new list" onPress={() => { setUpdate(!update) }} />
        </View>
    );
}

export default UserLists;